package cmd

import (
	"github.com/Bnei-Baruch/mdb/dal"
	"github.com/Bnei-Baruch/mdb/rest"
	"github.com/Bnei-Baruch/mdb/utils"
	"github.com/Bnei-Baruch/mdb/version"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stvp/rollbar"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/gin-contrib/cors.v1"

	"net/http"
    "math/rand"
    "time"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "MDB api server",
	Run: serverFn,
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func serverDefaults() {
	viper.SetDefault("server", map[string]interface{}{
		"bind-address": ":8080",
		"mode": "debug",
		"rollbar-token": "",
		"rollbar-environment": "development",
	})
}

func serverFn(cmd *cobra.Command, args []string) {
    rand.Seed(time.Now().UTC().UnixNano())
	serverDefaults()

	// Setup logging
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	// Setup rollbar
	rollbar.Token = viper.GetString("server.rollbar-token")
	rollbar.Environment = viper.GetString("server.rollbar-environment")
	rollbar.CodeVersion = version.Version

	// Setup gin
	gin.SetMode(viper.GetString("server.mode"))
	router := gin.New()

	var recovery gin.HandlerFunc
	if len(rollbar.Token) > 0 {
		recovery = utils.RollbarRecoveryMiddleware()
	} else {
		recovery = gin.Recovery()
	}

	router.Use(utils.MdbLoggerMiddleware(log.StandardLogger()),
		utils.ErrorHandlingMiddleware(),
        utils.GinBodyLogMiddleware,
		cors.Default(),
		recovery)

    dal.Init()
	router.POST("/operations/capture_start", CaptureStartHandler)
	router.POST("/operations/capture_stop", CaptureStopHandler)
	router.POST("/operations/demux", DemuxHandler)
	router.POST("/operations/send", SendHandler)

	collections := router.Group("collections")
	collections.POST("/", rest.CollectionsCreateHandler)

	router.GET("/recover", func(c *gin.Context) {
		panic("test recover")
	})

	log.Infoln("Running application")
	router.Run(viper.GetString("server.bind-address"))

	// This would be reasonable once we'll have graceful shutdown implemented
	//if len(rollbar.Token) > 0 {
	//	rollbar.Wait()
	//}
}

// Handlers
type HandlerFunc func() error

func Handle(c *gin.Context, h HandlerFunc) {
    if err := h(); err == nil {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": "error",
            "error": err.Error(),
        })
    }
}

func CaptureStartHandler(c *gin.Context) {
	var cs rest.CaptureStart
	if c.BindJSON(&cs) == nil {
        Handle(c, func() error { return dal.CaptureStart(cs) })
    }
}

func CaptureStopHandler(c *gin.Context) {
	var cs rest.CaptureStop
	if c.BindJSON(&cs) == nil {
        Handle(c, func() error { return dal.CaptureStop(cs) })
    }
}

func DemuxHandler(c *gin.Context) {
    var demux rest.Demux
	if c.BindJSON(&demux) == nil {
        Handle(c, func() error { return dal.Demux(demux) })
    }
}

func SendHandler(c *gin.Context) {
    var send rest.Send
	if c.BindJSON(&send) == nil {
        Handle(c, func() error { return dal.Send(send) })
    }
}
