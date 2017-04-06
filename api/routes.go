package api

import (
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/operations/capture_start", CaptureStartHandler)
	router.POST("/operations/capture_stop", CaptureStopHandler)
	router.POST("/operations/demux", DemuxHandler)
	router.POST("/operations/trim", TrimHandler)
	router.POST("/operations/send", SendHandler)
	router.POST("/operations/convert", ConvertHandler)
	router.POST("/operations/upload", UploadHandler)

	router.GET("/sources/hierarchy", SourcesHierarchyHandler)
	router.GET("/tags/hierarchy", TagsHierarchyHandler)

	// Serve the log file.
	admin := router.Group("admin")
	admin.StaticFile("/log", viper.GetString("server.log"))

	// Serve the auto generated docs.
	router.StaticFile("/docs", viper.GetString("server.docs"))

	rest := router.Group("rest")
	rest.GET("/collections/", CollectionsListHandler)
	rest.POST("/collections/:id/activate", CollectionActivateHandler)

	router.GET("/recover", func(c *gin.Context) {
		panic("test recover")
	})
}
