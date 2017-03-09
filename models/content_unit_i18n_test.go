package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testContentUnitI18ns(t *testing.T) {
	t.Parallel()

	query := ContentUnitI18ns(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testContentUnitI18nsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = contentUnitI18n.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := ContentUnitI18ns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContentUnitI18nsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ContentUnitI18ns(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := ContentUnitI18ns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContentUnitI18nsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ContentUnitI18nSlice{contentUnitI18n}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := ContentUnitI18ns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testContentUnitI18nsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ContentUnitI18nExists(tx, contentUnitI18n.ContentUnitID, contentUnitI18n.Language)
	if err != nil {
		t.Errorf("Unable to check if ContentUnitI18n exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ContentUnitI18nExistsG to return true, but got false.")
	}
}
func testContentUnitI18nsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx); err != nil {
		t.Error(err)
	}

	contentUnitI18nFound, err := FindContentUnitI18n(tx, contentUnitI18n.ContentUnitID, contentUnitI18n.Language)
	if err != nil {
		t.Error(err)
	}

	if contentUnitI18nFound == nil {
		t.Error("want a record, got nil")
	}
}
func testContentUnitI18nsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ContentUnitI18ns(tx).Bind(contentUnitI18n); err != nil {
		t.Error(err)
	}
}

func testContentUnitI18nsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := ContentUnitI18ns(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testContentUnitI18nsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentUnitI18nOne := &ContentUnitI18n{}
	contentUnitI18nTwo := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18nOne, contentUnitI18nDBTypes, false, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}
	if err = randomize.Struct(seed, contentUnitI18nTwo, contentUnitI18nDBTypes, false, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18nOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = contentUnitI18nTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := ContentUnitI18ns(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testContentUnitI18nsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	contentUnitI18nOne := &ContentUnitI18n{}
	contentUnitI18nTwo := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18nOne, contentUnitI18nDBTypes, false, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}
	if err = randomize.Struct(seed, contentUnitI18nTwo, contentUnitI18nDBTypes, false, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18nOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = contentUnitI18nTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ContentUnitI18ns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testContentUnitI18nsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ContentUnitI18ns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContentUnitI18nsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx, contentUnitI18nColumns...); err != nil {
		t.Error(err)
	}

	count, err := ContentUnitI18ns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContentUnitI18nToOneContentUnitUsingContentUnit(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local ContentUnitI18n
	var foreign ContentUnit

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, contentUnitI18nDBTypes, true, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, contentUnitDBTypes, true, contentUnitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnit struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ContentUnitID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ContentUnit(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ContentUnitI18nSlice{&local}
	if err = local.L.LoadContentUnit(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.ContentUnit == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ContentUnit = nil
	if err = local.L.LoadContentUnit(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ContentUnit == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testContentUnitI18nToOneUserUsingUser(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local ContentUnitI18n
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, contentUnitI18nDBTypes, true, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	local.UserID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.UserID.Int64 = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.User(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ContentUnitI18nSlice{&local}
	if err = local.L.LoadUser(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testContentUnitI18nToOneSetOpContentUnitUsingContentUnit(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a ContentUnitI18n
	var b, c ContentUnit

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contentUnitI18nDBTypes, false, strmangle.SetComplement(contentUnitI18nPrimaryKeyColumns, contentUnitI18nColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, contentUnitDBTypes, false, strmangle.SetComplement(contentUnitPrimaryKeyColumns, contentUnitColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, contentUnitDBTypes, false, strmangle.SetComplement(contentUnitPrimaryKeyColumns, contentUnitColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ContentUnit{&b, &c} {
		err = a.SetContentUnit(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ContentUnit != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ContentUnitI18ns[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ContentUnitID != x.ID {
			t.Error("foreign key was wrong value", a.ContentUnitID)
		}

		if exists, err := ContentUnitI18nExists(tx, a.ContentUnitID, a.Language); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testContentUnitI18nToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a ContentUnitI18n
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contentUnitI18nDBTypes, false, strmangle.SetComplement(contentUnitI18nPrimaryKeyColumns, contentUnitI18nColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ContentUnitI18ns[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID.Int64 != x.ID {
			t.Error("foreign key was wrong value", a.UserID.Int64)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID.Int64))
		reflect.Indirect(reflect.ValueOf(&a.UserID.Int64)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID.Int64 != x.ID {
			t.Error("foreign key was wrong value", a.UserID.Int64, x.ID)
		}
	}
}

func testContentUnitI18nToOneRemoveOpUserUsingUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a ContentUnitI18n
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contentUnitI18nDBTypes, false, strmangle.SetComplement(contentUnitI18nPrimaryKeyColumns, contentUnitI18nColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetUser(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveUser(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.User(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.User != nil {
		t.Error("R struct entry should be nil")
	}

	if a.UserID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ContentUnitI18ns) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testContentUnitI18nsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = contentUnitI18n.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testContentUnitI18nsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ContentUnitI18nSlice{contentUnitI18n}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testContentUnitI18nsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := ContentUnitI18ns(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	contentUnitI18nDBTypes = map[string]string{`ContentUnitID`: `bigint`, `CreatedAt`: `timestamp with time zone`, `Description`: `text`, `Language`: `character`, `Name`: `text`, `OriginalLanguage`: `character`, `UserID`: `bigint`}
	_                      = bytes.MinRead
)

func testContentUnitI18nsUpdate(t *testing.T) {
	t.Parallel()

	if len(contentUnitI18nColumns) == len(contentUnitI18nPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ContentUnitI18ns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true, contentUnitI18nColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	if err = contentUnitI18n.Update(tx); err != nil {
		t.Error(err)
	}
}

func testContentUnitI18nsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(contentUnitI18nColumns) == len(contentUnitI18nPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	contentUnitI18n := &ContentUnitI18n{}
	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ContentUnitI18ns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, contentUnitI18n, contentUnitI18nDBTypes, true, contentUnitI18nPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(contentUnitI18nColumns, contentUnitI18nPrimaryKeyColumns) {
		fields = contentUnitI18nColumns
	} else {
		fields = strmangle.SetComplement(
			contentUnitI18nColumns,
			contentUnitI18nPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(contentUnitI18n))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ContentUnitI18nSlice{contentUnitI18n}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testContentUnitI18nsUpsert(t *testing.T) {
	t.Parallel()

	if len(contentUnitI18nColumns) == len(contentUnitI18nPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	contentUnitI18n := ContentUnitI18n{}
	if err = randomize.Struct(seed, &contentUnitI18n, contentUnitI18nDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = contentUnitI18n.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert ContentUnitI18n: %s", err)
	}

	count, err := ContentUnitI18ns(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &contentUnitI18n, contentUnitI18nDBTypes, false, contentUnitI18nPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ContentUnitI18n struct: %s", err)
	}

	if err = contentUnitI18n.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert ContentUnitI18n: %s", err)
	}

	count, err = ContentUnitI18ns(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
