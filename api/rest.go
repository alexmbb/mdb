package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/nullbio/null.v6"

	"github.com/Bnei-Baruch/mdb/models"
	"github.com/Bnei-Baruch/mdb/utils"
)

const (
	DEFAULT_PAGE_SIZE = 50
	MAX_PAGE_SIZE     = 1000
)

func CollectionsListHandler(c *gin.Context) {
	var r CollectionsRequest
	if c.Bind(&r) != nil {
		return
	}

	resp, err := handleCollectionsList(boil.GetDB(), r)
	concludeRequest(c, resp, err)
}

func CollectionHandler(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("id"), 10, 0)
	if e != nil {
		NewBadRequestError(errors.Wrap(e, "id expects int64")).Abort(c)
		return
	}

	var err *HttpError
	var resp interface{}

	if c.Request.Method == http.MethodGet || c.Request.Method == "" {
		resp, err = handleGetCollection(boil.GetDB(), id)
	} else {
		if c.Request.Method == http.MethodPut {
			var cl Collection
			if c.Bind(&cl) != nil {
				return
			}

			cl.ID = id
			tx := mustBeginTx()
			resp, err = handleUpdateCollection(tx, &cl)
			mustConcludeTx(tx, err)
		}
	}

	concludeRequest(c, resp, err)
}

func CollectionContentUnitsHandler(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("id"), 10, 0)
	if e != nil {
		NewBadRequestError(errors.Wrap(e, "id expects int64")).Abort(c)
		return
	}

	resp, err := handleCollectionCCU(boil.GetDB(), id)
	concludeRequest(c, resp, err)
}

// Toggle the active flag of a single container
func CollectionActivateHandler(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("id"), 10, 0)
	if e != nil {
		NewBadRequestError(errors.Wrap(e, "id expects int64")).Abort(c)
		return
	}

	err := handleCollectionActivate(boil.GetDB(), id)
	concludeRequest(c, gin.H{"status": "ok"}, err)
}

func ContentUnitsListHandler(c *gin.Context) {
	var r ContentUnitsRequest
	if c.Bind(&r) != nil {
		return
	}

	resp, err := handleContentUnitsList(boil.GetDB(), r)
	concludeRequest(c, resp, err)
}

func ContentUnitHandler(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("id"), 10, 0)
	if e != nil {
		NewBadRequestError(errors.Wrap(e, "id expects int64")).Abort(c)
		return
	}

	var err *HttpError
	var resp interface{}

	if c.Request.Method == http.MethodGet || c.Request.Method == "" {
		resp, err = handleGetContentUnit(boil.GetDB(), id)
	} else {
		if c.Request.Method == http.MethodPut {
			var cu ContentUnit
			if c.Bind(&cu) != nil {
				return
			}

			cu.ID = id
			tx := mustBeginTx()
			resp, err = handleUpdateContentUnit(tx, &cu)
			mustConcludeTx(tx, err)
		}
	}

	concludeRequest(c, resp, err)
}

func ContentUnitFilesHandler(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("id"), 10, 0)
	if e != nil {
		NewBadRequestError(errors.Wrap(e, "id expects int64")).Abort(c)
		return
	}

	resp, err := handleContentUnitFiles(boil.GetDB(), id)
	concludeRequest(c, resp, err)
}

func ContentUnitCollectionsHandler(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("id"), 10, 0)
	if e != nil {
		NewBadRequestError(errors.Wrap(e, "id expects int64")).Abort(c)
		return
	}

	resp, err := handleContentUnitCCU(boil.GetDB(), id)
	concludeRequest(c, resp, err)
}

func FilesListHandler(c *gin.Context) {
	var r FilesRequest
	if c.Bind(&r) != nil {
		return
	}

	resp, err := handleFilesList(boil.GetDB(), r)
	concludeRequest(c, resp, err)
}

func FileHandler(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("id"), 10, 0)
	if e != nil {
		NewBadRequestError(errors.Wrap(e, "id expects int64")).Abort(c)
		return
	}

	var err *HttpError
	var resp interface{}

	if c.Request.Method == http.MethodGet || c.Request.Method == "" {
		resp, err = handleGetFile(boil.GetDB(), id)
	} else {
		if c.Request.Method == http.MethodPut {
			var f MFile
			if c.Bind(&f) != nil {
				return
			}

			f.ID = id
			tx := mustBeginTx()
			resp, err = handleUpdateFile(tx, &f)
			mustConcludeTx(tx, err)
		}
	}

	concludeRequest(c, resp, err)
}

func OperationsListHandler(c *gin.Context) {
	var r OperationsRequest
	if c.Bind(&r) != nil {
		return
	}

	resp, err := handleOperationsList(boil.GetDB(), r)
	concludeRequest(c, resp, err)
}

func OperationItemHandler(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("id"), 10, 0)
	if e != nil {
		NewBadRequestError(errors.Wrap(e, "id expects int64")).Abort(c)
		return
	}

	resp, err := handleOperationItem(boil.GetDB(), id)
	concludeRequest(c, resp, err)
}

func OperationFilesHandler(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("id"), 10, 0)
	if e != nil {
		NewBadRequestError(errors.Wrap(e, "id expects int64")).Abort(c)
		return
	}

	resp, err := handleOperationFiles(boil.GetDB(), id)
	concludeRequest(c, resp, err)
}

func TagsHandler(c *gin.Context) {
	var err *HttpError
	var resp interface{}

	if c.Request.Method == http.MethodGet || c.Request.Method == "" {
		var r TagsRequest
		if c.Bind(&r) != nil {
			return
		}
		resp, err = handleGetTags(boil.GetDB(), r)
	} else {
		if c.Request.Method == http.MethodPost {
			var t Tag
			if c.Bind(&t) != nil {
				return
			}

			for _, x := range t.I18n {
				if StdLang(x.Language) == LANG_UNKNOWN {
					NewBadRequestError(errors.Errorf("Unknown language %s", x.Language)).Abort(c)
					return
				}
			}

			tx := mustBeginTx()
			resp, err = handleCreateTag(tx, &t)
			mustConcludeTx(tx, err)
		}
	}

	concludeRequest(c, resp, err)
}

func TagHandler(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("id"), 10, 0)
	if e != nil {
		NewBadRequestError(errors.Wrap(e, "id expects int64")).Abort(c)
		return
	}

	var err *HttpError
	var resp interface{}

	if c.Request.Method == http.MethodGet || c.Request.Method == "" {
		resp, err = handleGetTag(boil.GetDB(), id)
	} else {
		if c.Request.Method == http.MethodPut {
			var t Tag
			if c.Bind(&t) != nil {
				return
			}

			t.ID = id
			tx := mustBeginTx()
			resp, err = handleUpdateTag(tx, &t)
			mustConcludeTx(tx, err)
		}
	}

	concludeRequest(c, resp, err)
}

func TagI18nHandler(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("id"), 10, 0)
	if e != nil {
		NewBadRequestError(errors.Wrap(e, "id expects int64")).Abort(c)
		return
	}

	var i18ns []*models.TagI18n
	if c.Bind(&i18ns) != nil {
		return
	}
	for _, x := range i18ns {
		if StdLang(x.Language) == LANG_UNKNOWN {
			NewBadRequestError(errors.Errorf("Unknown language %s", x.Language)).Abort(c)
			return
		}
	}

	tx := mustBeginTx()
	resp, err := handleUpdateTagI18n(tx, id, i18ns)
	mustConcludeTx(tx, err)
	concludeRequest(c, resp, err)
}

// Handlers Logic

func handleCollectionsList(exec boil.Executor, r CollectionsRequest) (*CollectionsResponse, *HttpError) {
	mods := make([]qm.QueryMod, 0)

	// filters
	if err := appendContentTypesFilterMods(&mods, r.ContentTypesFilter); err != nil {
		return nil, NewBadRequestError(err)
	}
	if err := appendDateRangeFilterMods(&mods, r.DateRangeFilter, "(properties->>'film_date')::date"); err != nil {
		return nil, NewBadRequestError(err)
	}
	if err := appendSecureFilterMods(&mods, r.SecureFilter); err != nil {
		return nil, NewBadRequestError(err)
	}
	appendPublishedFilterMods(&mods, r.PublishedFilter)

	// count query
	total, err := models.Collections(exec, mods...).Count()
	if err != nil {
		return nil, NewInternalError(err)
	}
	if total == 0 {
		return NewCollectionsResponse(), nil
	}

	// order, limit, offset
	if err = appendListMods(&mods, r.ListRequest); err != nil {
		return nil, NewBadRequestError(err)
	}

	// Eager loading
	mods = append(mods, qm.Load("CollectionI18ns"))

	// data query
	collections, err := models.Collections(exec, mods...).All()
	if err != nil {
		return nil, NewInternalError(err)
	}

	// i18n
	data := make([]*Collection, len(collections))
	for i, c := range collections {
		x := &Collection{Collection: *c}
		data[i] = x
		x.I18n = make(map[string]*models.CollectionI18n, len(c.R.CollectionI18ns))
		for _, i18n := range c.R.CollectionI18ns {
			x.I18n[i18n.Language] = i18n
		}
	}

	return &CollectionsResponse{
		ListResponse: ListResponse{Total: total},
		Collections:  data,
	}, nil
}

func handleGetCollection(exec boil.Executor, id int64) (*Collection, *HttpError) {
	collection, err := models.Collections(exec,
		qm.Where("id = ?", id),
		qm.Load("CollectionI18ns")).
		One()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewNotFoundError()
		} else {
			return nil, NewInternalError(err)
		}
	}

	// i18n
	x := &Collection{Collection: *collection}
	x.I18n = make(map[string]*models.CollectionI18n, len(collection.R.CollectionI18ns))
	for _, i18n := range collection.R.CollectionI18ns {
		x.I18n[i18n.Language] = i18n
	}

	return x, nil
}

func handleUpdateCollection(exec boil.Executor, c *Collection) (*Collection, *HttpError) {
	collection, err := models.FindCollection(exec, c.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewNotFoundError()
		} else {
			return nil, NewInternalError(err)
		}
	}

	collection.Secure = c.Secure
	err = collection.Update(exec, "secure")
	if err != nil {
		return nil, NewInternalError(err)
	}

	return handleGetCollection(exec, c.ID)
}

func handleCollectionActivate(exec boil.Executor, id int64) *HttpError {
	collection, err := models.FindCollection(exec, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return NewNotFoundError()
		} else {
			return NewInternalError(err)
		}
	}

	var props = make(map[string]interface{})
	if collection.Properties.Valid {
		collection.Properties.Unmarshal(&props)
	}
	active, ok := props["active"]
	if ok {
		b, _ := active.(bool)
		props["active"] = !b
	} else {
		props["active"] = false
	}

	pbytes, err := json.Marshal(props)
	if err != nil {
		return NewInternalError(err)
	}
	collection.Properties = null.JSONFrom(pbytes)
	err = collection.Update(exec, "properties")
	if err != nil {
		return NewInternalError(err)
	}

	return nil
}

func handleCollectionCCU(exec boil.Executor, id int64) ([]*CollectionContentUnit, *HttpError) {
	ok, err := models.CollectionExists(exec, id)
	if err != nil {
		return nil, NewInternalError(err)
	}
	if !ok {
		return nil, NewNotFoundError()
	}

	ccus, err := models.CollectionsContentUnits(exec, qm.Where("collection_id = ?", id)).All()
	if err != nil {
		return nil, NewInternalError(err)
	} else if len(ccus) == 0 {
		return make([]*CollectionContentUnit, 0), nil
	}

	ids := make([]int64, len(ccus))
	for i, ccu := range ccus {
		ids[i] = ccu.ContentUnitID
	}
	cus, err := models.ContentUnits(exec,
		qm.WhereIn("id in ?", utils.ConvertArgsInt64(ids)...),
		qm.Load("ContentUnitI18ns")).
		All()
	if err != nil {
		return nil, NewInternalError(err)
	}

	cusById := make(map[int64]*ContentUnit, len(cus))
	for _, cu := range cus {
		x := ContentUnit{ContentUnit: *cu}
		x.I18n = make(map[string]*models.ContentUnitI18n, len(cu.R.ContentUnitI18ns))
		for _, i18n := range cu.R.ContentUnitI18ns {
			x.I18n[i18n.Language] = i18n
		}
		cusById[x.ID] = &x
	}

	data := make([]*CollectionContentUnit, len(ccus))
	for i, ccu := range ccus {
		data[i] = &CollectionContentUnit{
			Name:        ccu.Name,
			ContentUnit: cusById[ccu.ContentUnitID],
		}
	}

	return data, nil
}

func handleContentUnitsList(exec boil.Executor, r ContentUnitsRequest) (*ContentUnitsResponse, *HttpError) {
	mods := make([]qm.QueryMod, 0)

	// filters
	if err := appendContentTypesFilterMods(&mods, r.ContentTypesFilter); err != nil {
		return nil, NewBadRequestError(err)
	}
	if err := appendDateRangeFilterMods(&mods, r.DateRangeFilter, "(properties->>'film_date')::date"); err != nil {
		return nil, NewBadRequestError(err)
	}
	if err := appendSourcesFilterMods(exec, &mods, r.SourcesFilter); err != nil {
		if e, ok := err.(*HttpError); ok {
			return nil, e
		} else {
			NewInternalError(err)
		}
	}
	if err := appendTagsFilterMods(exec, &mods, r.TagsFilter); err != nil {
		return nil, NewInternalError(err)
	}
	if err := appendSecureFilterMods(&mods, r.SecureFilter); err != nil {
		return nil, NewBadRequestError(err)
	}
	appendPublishedFilterMods(&mods, r.PublishedFilter)

	// count query
	total, err := models.ContentUnits(exec, mods...).Count()
	if err != nil {
		return nil, NewInternalError(err)
	}
	if total == 0 {
		return NewContentUnitsResponse(), nil
	}

	// order, limit, offset
	if err = appendListMods(&mods, r.ListRequest); err != nil {
		return nil, NewBadRequestError(err)
	}

	// Eager loading
	mods = append(mods, qm.Load("ContentUnitI18ns"))

	// data query
	units, err := models.ContentUnits(exec, mods...).All()
	if err != nil {
		return nil, NewInternalError(err)
	}

	// i18n
	data := make([]*ContentUnit, len(units))
	for i, cu := range units {
		x := &ContentUnit{ContentUnit: *cu}
		data[i] = x
		x.I18n = make(map[string]*models.ContentUnitI18n, len(cu.R.ContentUnitI18ns))
		for _, i18n := range cu.R.ContentUnitI18ns {
			x.I18n[i18n.Language] = i18n
		}
	}

	return &ContentUnitsResponse{
		ListResponse: ListResponse{Total: total},
		ContentUnits: data,
	}, nil
}

func handleGetContentUnit(exec boil.Executor, id int64) (*ContentUnit, *HttpError) {
	unit, err := models.ContentUnits(exec,
		qm.Where("id = ?", id),
		qm.Load("ContentUnitI18ns")).
		One()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewNotFoundError()
		} else {
			return nil, NewInternalError(err)
		}
	}

	// i18n
	x := &ContentUnit{ContentUnit: *unit}
	x.I18n = make(map[string]*models.ContentUnitI18n, len(unit.R.ContentUnitI18ns))
	for _, i18n := range unit.R.ContentUnitI18ns {
		x.I18n[i18n.Language] = i18n
	}

	return x, nil
}


func handleUpdateContentUnit(exec boil.Executor, cu *ContentUnit) (*ContentUnit, *HttpError) {
	unit, err := models.FindContentUnit(exec, cu.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewNotFoundError()
		} else {
			return nil, NewInternalError(err)
		}
	}

	unit.Secure = cu.Secure
	err = unit.Update(exec, "secure")
	if err != nil {
		return nil, NewInternalError(err)
	}

	return handleGetContentUnit(exec, cu.ID)
}

func handleContentUnitFiles(exec boil.Executor, id int64) ([]*MFile, *HttpError) {
	ok, err := models.ContentUnitExists(exec, id)
	if err != nil {
		return nil, NewInternalError(err)
	}
	if !ok {
		return nil, NewNotFoundError()
	}

	files, err := models.Files(exec, qm.Where("content_unit_id = ?", id)).All()
	if err != nil {
		return nil, NewInternalError(err)
	}

	data := make([]*MFile, len(files))
	for i, f := range files {
		data[i] = NewMFile(f)
	}

	return data, nil
}

func handleContentUnitCCU(exec boil.Executor, id int64) ([]*CollectionContentUnit, *HttpError) {
	ok, err := models.ContentUnitExists(exec, id)
	if err != nil {
		return nil, NewInternalError(err)
	}
	if !ok {
		return nil, NewNotFoundError()
	}

	ccus, err := models.CollectionsContentUnits(exec, qm.Where("content_unit_id = ?", id)).All()
	if err != nil {
		return nil, NewInternalError(err)
	} else if len(ccus) == 0 {
		return make([]*CollectionContentUnit, 0), nil
	}

	ids := make([]int64, len(ccus))
	for i, ccu := range ccus {
		ids[i] = ccu.CollectionID
	}
	cs, err := models.Collections(exec,
		qm.WhereIn("id in ?", utils.ConvertArgsInt64(ids)...),
		qm.Load("CollectionI18ns")).
		All()
	if err != nil {
		return nil, NewInternalError(err)
	}

	csById := make(map[int64]*Collection, len(cs))
	for _, c := range cs {
		x := Collection{Collection: *c}
		x.I18n = make(map[string]*models.CollectionI18n, len(c.R.CollectionI18ns))
		for _, i18n := range c.R.CollectionI18ns {
			x.I18n[i18n.Language] = i18n
		}
		csById[x.ID] = &x
	}

	data := make([]*CollectionContentUnit, len(ccus))
	for i, ccu := range ccus {
		data[i] = &CollectionContentUnit{
			Name:       ccu.Name,
			Collection: csById[ccu.CollectionID],
		}
	}

	return data, nil
}

func handleFilesList(exec boil.Executor, r FilesRequest) (*FilesResponse, *HttpError) {

	mods := make([]qm.QueryMod, 0)

	// filters
	if err := appendDateRangeFilterMods(&mods, r.DateRangeFilter, "file_created_at"); err != nil {
		return nil, NewBadRequestError(err)
	}
	if err := appendSecureFilterMods(&mods, r.SecureFilter); err != nil {
		return nil, NewBadRequestError(err)
	}
	appendPublishedFilterMods(&mods, r.PublishedFilter)
	if r.Query != "" {
		mods = append(mods, qm.Where("name ~ ?", r.Query),
			qm.Or("uid ~ ?", r.Query),
			qm.Or("id::TEXT ~ ?", r.Query))
	}

	// count query
	total, err := models.Files(exec, mods...).Count()
	if err != nil {
		return nil, NewInternalError(err)
	}
	if total == 0 {
		return NewFilesResponse(), nil
	}

	// order, limit, offset
	if err = appendListMods(&mods, r.ListRequest); err != nil {
		return nil, NewBadRequestError(err)
	}

	// data query
	files, err := models.Files(exec, mods...).All()
	if err != nil {
		return nil, NewInternalError(err)
	}

	data := make([]*MFile, len(files))
	for i, f := range files {
		data[i] = NewMFile(f)
	}

	return &FilesResponse{
		ListResponse: ListResponse{Total: total},
		Files:        data,
	}, nil
}

func handleGetFile(exec boil.Executor, id int64) (*MFile, *HttpError) {
	file, err := models.FindFile(exec, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewNotFoundError()
		} else {
			return nil, NewInternalError(err)
		}
	}

	return NewMFile(file), nil
}

func handleUpdateFile(exec boil.Executor, f *MFile) (*MFile, *HttpError) {
	file, err := models.FindFile(exec, f.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewNotFoundError()
		} else {
			return nil, NewInternalError(err)
		}
	}

	file.Secure = f.Secure
	err = file.Update(exec, "secure")
	if err != nil {
		return nil, NewInternalError(err)
	}

	return handleGetFile(exec, f.ID)
}

func handleOperationsList(exec boil.Executor, r OperationsRequest) (*OperationsResponse, *HttpError) {

	mods := make([]qm.QueryMod, 0)

	// filters
	if err := appendDateRangeFilterMods(&mods, r.DateRangeFilter, "created_at"); err != nil {
		return nil, NewBadRequestError(err)
	}
	if err := appendOperationTypesFilterMods(&mods, r.OperationTypesFilter); err != nil {
		return nil, NewBadRequestError(err)
	}

	// count query
	total, err := models.Operations(exec, mods...).Count()
	if err != nil {
		return nil, NewInternalError(err)
	}
	if total == 0 {
		return NewOperationsResponse(), nil
	}

	// order, limit, offset
	if err = appendListMods(&mods, r.ListRequest); err != nil {
		return nil, NewBadRequestError(err)
	}

	// data query
	data, err := models.Operations(exec, mods...).All()
	if err != nil {
		return nil, NewInternalError(err)
	}

	return &OperationsResponse{
		ListResponse: ListResponse{Total: total},
		Operations:   data,
	}, nil
}

func handleOperationItem(exec boil.Executor, id int64) (*models.Operation, *HttpError) {
	operation, err := models.FindOperation(exec, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewNotFoundError()
		} else {
			return nil, NewInternalError(err)
		}
	}

	return operation, nil
}

func handleOperationFiles(exec boil.Executor, id int64) ([]*MFile, *HttpError) {
	ok, err := models.OperationExists(exec, id)
	if err != nil {
		return nil, NewInternalError(err)
	}
	if !ok {
		return nil, NewNotFoundError()
	}

	files, err := models.Files(exec,
		qm.InnerJoin("files_operations fo on fo.file_id=id and fo.operation_id = ?", id)).
		All()
	if err != nil {
		return nil, NewInternalError(err)
	}

	data := make([]*MFile, len(files))
	for i, f := range files {
		data[i] = NewMFile(f)
	}

	return data, nil
}

func handleGetTags(exec boil.Executor, r TagsRequest) (*TagsResponse, *HttpError) {
	mods := make([]qm.QueryMod, 0)

	// count query
	total, err := models.Tags(exec, mods...).Count()
	if err != nil {
		return nil, NewInternalError(err)
	}
	if total == 0 {
		return NewTagsResponse(), nil
	}

	// order, limit, offset
	if err = appendListMods(&mods, r.ListRequest); err != nil {
		return nil, NewBadRequestError(err)
	}

	// Eager loading
	mods = append(mods, qm.Load("TagI18ns"))

	// data query
	tags, err := models.Tags(exec, mods...).All()
	if err != nil {
		return nil, NewInternalError(err)
	}

	// i18n
	data := make([]*Tag, len(tags))
	for i, t := range tags {
		x := &Tag{Tag: *t}
		data[i] = x
		x.I18n = make(map[string]*models.TagI18n, len(t.R.TagI18ns))
		for _, i18n := range t.R.TagI18ns {
			x.I18n[i18n.Language] = i18n
		}
	}

	return &TagsResponse{
		ListResponse: ListResponse{Total: total},
		Tags:         data,
	}, nil
}

func handleCreateTag(exec boil.Executor, t *Tag) (*Tag, *HttpError) {
	// make sure parent tag exists if given
	if t.ParentID.Valid {
		ok, err := models.Tags(exec, qm.Where("id = ?", t.ParentID.Int64)).Exists()
		if err != nil {
			return nil, NewInternalError(err)
		}
		if !ok {
			return nil, NewBadRequestError(errors.Errorf("Unknown parent tag %d", t.ParentID.Int64))
		}
	}

	// save tag to DB
	var uid string
	for {
		uid = utils.GenerateUID(8)
		exists, err := models.ContentUnits(exec, qm.Where("uid = ?", uid)).Exists()
		if err != nil {
			return nil, NewInternalError(err)
		}
		if !exists {
			break
		}
	}

	t.UID = uid
	err := t.Tag.Insert(exec)
	if err != nil {
		return nil, NewInternalError(err)
	}

	// save i18n
	for _, v := range t.I18n {
		err := t.AddTagI18ns(exec, true, v)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	return handleGetTag(exec, t.ID)
}

func handleGetTag(exec boil.Executor, id int64) (*Tag, *HttpError) {
	tag, err := models.Tags(exec,
		qm.Where("id = ?", id),
		qm.Load("TagI18ns")).
		One()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewNotFoundError()
		} else {
			return nil, NewInternalError(err)
		}
	}

	// i18n
	x := &Tag{Tag: *tag}
	x.I18n = make(map[string]*models.TagI18n, len(tag.R.TagI18ns))
	for _, i18n := range tag.R.TagI18ns {
		x.I18n[i18n.Language] = i18n
	}

	return x, nil
}

func handleUpdateTag(exec boil.Executor, t *Tag) (*Tag, *HttpError) {
	tag, err := models.FindTag(exec, t.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewNotFoundError()
		} else {
			return nil, NewInternalError(err)
		}
	}

	tag.Pattern = t.Pattern
	tag.Description = t.Description
	err = t.Update(exec, "pattern", "description")
	if err != nil {
		return nil, NewInternalError(err)
	}

	return handleGetTag(exec, t.ID)
}

func handleUpdateTagI18n(exec boil.Executor, id int64, i18ns []*models.TagI18n) (*Tag, *HttpError) {
	tag, err := handleGetTag(exec, id)
	if err != nil {
		return nil, err
	}

	// Upsert all new i18ns
	nI18n := make(map[string]*models.TagI18n, len(i18ns))
	for _, i18n := range i18ns {
		i18n.TagID = id
		nI18n[i18n.Language] = i18n
		err := i18n.Upsert(exec, true, []string{"tag_id", "language"}, []string{"label"})
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	// Delete old i18ns not in new i18ns
	for k, v := range tag.I18n {
		if _, ok := nI18n[k]; !ok {
			err := v.Delete(exec)
			if err != nil {
				return nil, NewInternalError(err)
			}
		}
	}

	return handleGetTag(exec, id)
}

// Query Helpers

func appendListMods(mods *[]qm.QueryMod, r ListRequest) error {
	if r.OrderBy == "" {
		*mods = append(*mods, qm.OrderBy("id desc"))
	} else {
		*mods = append(*mods, qm.OrderBy(r.OrderBy))
	}

	var limit, offset int

	if r.StartIndex == 0 {
		// pagination style
		if r.PageSize == 0 {
			limit = DEFAULT_PAGE_SIZE
		} else {
			limit = utils.Min(r.PageSize, MAX_PAGE_SIZE)
		}
		if r.PageNumber > 1 {
			offset = (r.PageNumber - 1) * limit
		}
	} else {
		// start & stop index style for "infinite" lists
		offset = r.StartIndex - 1
		if r.StopIndex == 0 {
			limit = MAX_PAGE_SIZE
		} else if r.StopIndex < r.StartIndex {
			return errors.Errorf("Invalid range [%d-%d]", r.StartIndex, r.StopIndex)
		} else {
			limit = r.StopIndex - r.StartIndex + 1
		}
	}

	*mods = append(*mods, qm.Limit(limit))
	if offset != 0 {
		*mods = append(*mods, qm.Offset(offset))
	}

	return nil
}

func appendDateRangeFilterMods(mods *[]qm.QueryMod, f DateRangeFilter, field string) error {
	s, e, err := f.Range()
	if err != nil {
		return err
	}

	if f.StartDate != "" && f.EndDate != "" && e.Before(s) {
		return errors.New("Invalid date range")
	}

	if field == "" {
		field = "created_at"
	}

	if f.StartDate != "" {
		*mods = append(*mods, qm.Where(fmt.Sprintf("%s >= ?", field), s))
	}
	if f.EndDate != "" {
		*mods = append(*mods, qm.Where(fmt.Sprintf("%s <= ?", field), e))
	}

	return nil
}

func appendContentTypesFilterMods(mods *[]qm.QueryMod, f ContentTypesFilter) error {
	if utils.IsEmpty(f.ContentTypes) {
		return nil
	}

	a := make([]interface{}, len(f.ContentTypes))
	for i, x := range f.ContentTypes {
		ct, ok := CONTENT_TYPE_REGISTRY.ByName[strings.ToUpper(x)]
		if ok {
			a[i] = ct.ID
		} else {
			return errors.Errorf("Unknown content type: %s", x)
		}
	}

	*mods = append(*mods, qm.WhereIn("type_id in ?", a...))

	return nil
}

func appendSecureFilterMods(mods *[]qm.QueryMod, f SecureFilter) error {
	if len(f.Levels) == 0 {
		return nil
	}

	a := make([]interface{}, len(f.Levels))
	for i, x := range f.Levels {
		if x == SEC_PUBLIC || x == SEC_SENSITIVE || x == SEC_PRIVATE {
			a[i] = x
		} else {
			return errors.Errorf("Unknown security level: %d", x)
		}
	}

	*mods = append(*mods, qm.WhereIn("secure in ?", a...))

	return nil
}

func appendPublishedFilterMods(mods *[]qm.QueryMod, f PublishedFilter) {
	var val null.Bool
	val.UnmarshalText([]byte(f.Published))
	if val.Valid {
		*mods = append(*mods, qm.Where("published = ?", val.Bool))
	}
}

func appendOperationTypesFilterMods(mods *[]qm.QueryMod, f OperationTypesFilter) error {
	if utils.IsEmpty(f.OperationTypes) {
		return nil
	}

	a := make([]interface{}, len(f.OperationTypes))
	for i, x := range f.OperationTypes {
		ot, ok := OPERATION_TYPE_REGISTRY.ByName[strings.ToLower(x)]
		if ok {
			a[i] = ot.ID
		} else {
			return errors.Errorf("Unknown operation type: %s", x)
		}
	}

	*mods = append(*mods, qm.WhereIn("type_id in ?", a...))

	return nil
}

func appendSourcesFilterMods(exec boil.Executor, mods *[]qm.QueryMod, f SourcesFilter) error {
	// slice of all source ids we want
	source_ids := make([]int64, 0)

	// fetch source ids by authors
	if !utils.IsEmpty(f.Authors) {
		for _, x := range f.Authors {
			if _, ok := AUTHOR_REGISTRY.ByCode[strings.ToLower(x)]; !ok {
				return NewBadRequestError(errors.Errorf("Unknown author: %s", x))
			}
		}

		var ids pq.Int64Array
		q := `SELECT array_agg(DISTINCT "as".source_id)
		      FROM authors a INNER JOIN authors_sources "as" ON a.id = "as".author_id
		      WHERE a.code = ANY($1)`
		err := queries.Raw(exec, q, pq.Array(f.Authors)).QueryRow().Scan(&ids)
		if err != nil {
			return err
		}
		source_ids = append(source_ids, ids...)
	}

	// blend in requested sources
	source_ids = append(source_ids, f.Sources...)

	if len(source_ids) == 0 {
		return nil
	}

	// find all nested source_ids
	q := `WITH RECURSIVE rec_sources AS (
		  SELECT s.id FROM sources s WHERE s.id = ANY($1)
		  UNION
		  SELECT s.id FROM sources s INNER JOIN rec_sources rs ON s.parent_id = rs.id
	      )
	      SELECT array_agg(distinct id) FROM rec_sources`
	var ids pq.Int64Array
	err := queries.Raw(exec, q, pq.Array(source_ids)).QueryRow().Scan(&ids)
	if err != nil {
		return err
	}

	if ids == nil || len(ids) == 0 {
		*mods = append(*mods, qm.Where("id < 0")) // so results would be empty
	} else {
		*mods = append(*mods,
			qm.InnerJoin("content_units_sources cus ON id = cus.content_unit_id"),
			qm.WhereIn("cus.source_id in ?", utils.ConvertArgsInt64(ids)...))
	}

	return nil
}

func appendTagsFilterMods(exec boil.Executor, mods *[]qm.QueryMod, f TagsFilter) error {
	if len(f.Tags) == 0 {
		return nil
	}

	// find all nested tag_ids
	q := `WITH RECURSIVE rec_tags AS (
	        SELECT t.id FROM tags t WHERE t.id = ANY($1)
	        UNION
	        SELECT t.id FROM tags t INNER JOIN rec_tags rt ON t.parent_id = rt.id
	      )
	      SELECT array_agg(distinct id) FROM rec_tags`
	var ids pq.Int64Array
	err := queries.Raw(exec, q, pq.Array(f.Tags)).QueryRow().Scan(&ids)
	if err != nil {
		return err
	}

	if ids == nil || len(ids) == 0 {
		*mods = append(*mods, qm.Where("id < 0")) // so results would be empty
	} else {
		*mods = append(*mods,
			qm.InnerJoin("content_units_tags cut ON id = cut.content_unit_id"),
			qm.WhereIn("cut.tag_id in ?", utils.ConvertArgsInt64(ids)...))
	}

	return nil
}

// mustBeginTx begins a transaction, panics on error.
func mustBeginTx() boil.Transactor {
	tx, ex := boil.Begin()
	utils.Must(ex)
	return tx
}

// mustConcludeTx commits or rollback the given transaction according to given error.
// Panics if Commit() or Rollback() fails.
func mustConcludeTx(tx boil.Transactor, err *HttpError) {
	if err == nil {
		utils.Must(tx.Commit())
	} else {
		utils.Must(tx.Rollback())
	}
}

// concludeRequest responds with JSON of given response or aborts the request with the given error.
func concludeRequest(c *gin.Context, resp interface{}, err *HttpError) {
	if err == nil {
		c.JSON(http.StatusOK, resp)
	} else {
		err.Abort(c)
	}
}
