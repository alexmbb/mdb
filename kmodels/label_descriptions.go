package kmodels

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v6"
)

// LabelDescription is an object representing the database table.
type LabelDescription struct {
	ID        int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	LabelID   null.Int    `boil:"label_id" json:"label_id,omitempty" toml:"label_id" yaml:"label_id,omitempty"`
	Text      null.String `boil:"text" json:"text,omitempty" toml:"text" yaml:"text,omitempty"`
	Lang      null.String `boil:"lang" json:"lang,omitempty" toml:"lang" yaml:"lang,omitempty"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *labelDescriptionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L labelDescriptionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// labelDescriptionR is where relationships are stored.
type labelDescriptionR struct {
}

// labelDescriptionL is where Load methods for each relationship are stored.
type labelDescriptionL struct{}

var (
	labelDescriptionColumns               = []string{"id", "label_id", "text", "lang", "created_at", "updated_at"}
	labelDescriptionColumnsWithoutDefault = []string{"label_id", "text", "created_at", "updated_at"}
	labelDescriptionColumnsWithDefault    = []string{"id", "lang"}
	labelDescriptionPrimaryKeyColumns     = []string{"id"}
)

type (
	// LabelDescriptionSlice is an alias for a slice of pointers to LabelDescription.
	// This should generally be used opposed to []LabelDescription.
	LabelDescriptionSlice []*LabelDescription

	labelDescriptionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	labelDescriptionType                 = reflect.TypeOf(&LabelDescription{})
	labelDescriptionMapping              = queries.MakeStructMapping(labelDescriptionType)
	labelDescriptionPrimaryKeyMapping, _ = queries.BindMapping(labelDescriptionType, labelDescriptionMapping, labelDescriptionPrimaryKeyColumns)
	labelDescriptionInsertCacheMut       sync.RWMutex
	labelDescriptionInsertCache          = make(map[string]insertCache)
	labelDescriptionUpdateCacheMut       sync.RWMutex
	labelDescriptionUpdateCache          = make(map[string]updateCache)
	labelDescriptionUpsertCacheMut       sync.RWMutex
	labelDescriptionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single labelDescription record from the query, and panics on error.
func (q labelDescriptionQuery) OneP() *LabelDescription {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single labelDescription record from the query.
func (q labelDescriptionQuery) One() (*LabelDescription, error) {
	o := &LabelDescription{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "kmodels: failed to execute a one query for label_descriptions")
	}

	return o, nil
}

// AllP returns all LabelDescription records from the query, and panics on error.
func (q labelDescriptionQuery) AllP() LabelDescriptionSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all LabelDescription records from the query.
func (q labelDescriptionQuery) All() (LabelDescriptionSlice, error) {
	var o LabelDescriptionSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "kmodels: failed to assign all query results to LabelDescription slice")
	}

	return o, nil
}

// CountP returns the count of all LabelDescription records in the query, and panics on error.
func (q labelDescriptionQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all LabelDescription records in the query.
func (q labelDescriptionQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "kmodels: failed to count label_descriptions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q labelDescriptionQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q labelDescriptionQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "kmodels: failed to check if label_descriptions exists")
	}

	return count > 0, nil
}

// LabelDescriptionsG retrieves all records.
func LabelDescriptionsG(mods ...qm.QueryMod) labelDescriptionQuery {
	return LabelDescriptions(boil.GetDB(), mods...)
}

// LabelDescriptions retrieves all the records using an executor.
func LabelDescriptions(exec boil.Executor, mods ...qm.QueryMod) labelDescriptionQuery {
	mods = append(mods, qm.From("\"label_descriptions\""))
	return labelDescriptionQuery{NewQuery(exec, mods...)}
}

// FindLabelDescriptionG retrieves a single record by ID.
func FindLabelDescriptionG(id int, selectCols ...string) (*LabelDescription, error) {
	return FindLabelDescription(boil.GetDB(), id, selectCols...)
}

// FindLabelDescriptionGP retrieves a single record by ID, and panics on error.
func FindLabelDescriptionGP(id int, selectCols ...string) *LabelDescription {
	retobj, err := FindLabelDescription(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindLabelDescription retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLabelDescription(exec boil.Executor, id int, selectCols ...string) (*LabelDescription, error) {
	labelDescriptionObj := &LabelDescription{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"label_descriptions\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(labelDescriptionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "kmodels: unable to select from label_descriptions")
	}

	return labelDescriptionObj, nil
}

// FindLabelDescriptionP retrieves a single record by ID with an executor, and panics on error.
func FindLabelDescriptionP(exec boil.Executor, id int, selectCols ...string) *LabelDescription {
	retobj, err := FindLabelDescription(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *LabelDescription) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *LabelDescription) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *LabelDescription) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *LabelDescription) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("kmodels: no label_descriptions provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(labelDescriptionColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	labelDescriptionInsertCacheMut.RLock()
	cache, cached := labelDescriptionInsertCache[key]
	labelDescriptionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			labelDescriptionColumns,
			labelDescriptionColumnsWithDefault,
			labelDescriptionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(labelDescriptionType, labelDescriptionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(labelDescriptionType, labelDescriptionMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"label_descriptions\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "kmodels: unable to insert into label_descriptions")
	}

	if !cached {
		labelDescriptionInsertCacheMut.Lock()
		labelDescriptionInsertCache[key] = cache
		labelDescriptionInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single LabelDescription record. See Update for
// whitelist behavior description.
func (o *LabelDescription) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single LabelDescription record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *LabelDescription) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the LabelDescription, and panics on error.
// See Update for whitelist behavior description.
func (o *LabelDescription) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the LabelDescription.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *LabelDescription) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	key := makeCacheKey(whitelist, nil)
	labelDescriptionUpdateCacheMut.RLock()
	cache, cached := labelDescriptionUpdateCache[key]
	labelDescriptionUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(labelDescriptionColumns, labelDescriptionPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("kmodels: unable to update label_descriptions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"label_descriptions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, labelDescriptionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(labelDescriptionType, labelDescriptionMapping, append(wl, labelDescriptionPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "kmodels: unable to update label_descriptions row")
	}

	if !cached {
		labelDescriptionUpdateCacheMut.Lock()
		labelDescriptionUpdateCache[key] = cache
		labelDescriptionUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q labelDescriptionQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q labelDescriptionQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "kmodels: unable to update all for label_descriptions")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o LabelDescriptionSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o LabelDescriptionSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o LabelDescriptionSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LabelDescriptionSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("kmodels: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), labelDescriptionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"label_descriptions\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(labelDescriptionPrimaryKeyColumns), len(colNames)+1, len(labelDescriptionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "kmodels: unable to update all in labelDescription slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *LabelDescription) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *LabelDescription) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *LabelDescription) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *LabelDescription) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("kmodels: no label_descriptions provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	nzDefaults := queries.NonZeroDefaultSet(labelDescriptionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	labelDescriptionUpsertCacheMut.RLock()
	cache, cached := labelDescriptionUpsertCache[key]
	labelDescriptionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			labelDescriptionColumns,
			labelDescriptionColumnsWithDefault,
			labelDescriptionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			labelDescriptionColumns,
			labelDescriptionPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("kmodels: unable to upsert label_descriptions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(labelDescriptionPrimaryKeyColumns))
			copy(conflict, labelDescriptionPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"label_descriptions\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(labelDescriptionType, labelDescriptionMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(labelDescriptionType, labelDescriptionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "kmodels: unable to upsert label_descriptions")
	}

	if !cached {
		labelDescriptionUpsertCacheMut.Lock()
		labelDescriptionUpsertCache[key] = cache
		labelDescriptionUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single LabelDescription record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *LabelDescription) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single LabelDescription record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *LabelDescription) DeleteG() error {
	if o == nil {
		return errors.New("kmodels: no LabelDescription provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single LabelDescription record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *LabelDescription) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single LabelDescription record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LabelDescription) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("kmodels: no LabelDescription provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), labelDescriptionPrimaryKeyMapping)
	sql := "DELETE FROM \"label_descriptions\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "kmodels: unable to delete from label_descriptions")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q labelDescriptionQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q labelDescriptionQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("kmodels: no labelDescriptionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "kmodels: unable to delete all from label_descriptions")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o LabelDescriptionSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o LabelDescriptionSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("kmodels: no LabelDescription slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o LabelDescriptionSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LabelDescriptionSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("kmodels: no LabelDescription slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), labelDescriptionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"label_descriptions\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, labelDescriptionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(labelDescriptionPrimaryKeyColumns), 1, len(labelDescriptionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "kmodels: unable to delete all from labelDescription slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *LabelDescription) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *LabelDescription) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *LabelDescription) ReloadG() error {
	if o == nil {
		return errors.New("kmodels: no LabelDescription provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *LabelDescription) Reload(exec boil.Executor) error {
	ret, err := FindLabelDescription(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *LabelDescriptionSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *LabelDescriptionSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LabelDescriptionSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("kmodels: empty LabelDescriptionSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LabelDescriptionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	labelDescriptions := LabelDescriptionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), labelDescriptionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"label_descriptions\".* FROM \"label_descriptions\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, labelDescriptionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(labelDescriptionPrimaryKeyColumns), 1, len(labelDescriptionPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&labelDescriptions)
	if err != nil {
		return errors.Wrap(err, "kmodels: unable to reload all in LabelDescriptionSlice")
	}

	*o = labelDescriptions

	return nil
}

// LabelDescriptionExists checks if the LabelDescription row exists.
func LabelDescriptionExists(exec boil.Executor, id int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"label_descriptions\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "kmodels: unable to check if label_descriptions exists")
	}

	return exists, nil
}

// LabelDescriptionExistsG checks if the LabelDescription row exists.
func LabelDescriptionExistsG(id int) (bool, error) {
	return LabelDescriptionExists(boil.GetDB(), id)
}

// LabelDescriptionExistsGP checks if the LabelDescription row exists. Panics on error.
func LabelDescriptionExistsGP(id int) bool {
	e, err := LabelDescriptionExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// LabelDescriptionExistsP checks if the LabelDescription row exists. Panics on error.
func LabelDescriptionExistsP(exec boil.Executor, id int) bool {
	e, err := LabelDescriptionExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}