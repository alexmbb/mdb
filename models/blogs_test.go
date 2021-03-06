// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testBlogs(t *testing.T) {
	t.Parallel()

	query := Blogs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testBlogsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = blog.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Blogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlogsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Blogs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Blogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlogsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BlogSlice{blog}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Blogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testBlogsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := BlogExists(tx, blog.ID)
	if err != nil {
		t.Errorf("Unable to check if Blog exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BlogExistsG to return true, but got false.")
	}
}
func testBlogsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx); err != nil {
		t.Error(err)
	}

	blogFound, err := FindBlog(tx, blog.ID)
	if err != nil {
		t.Error(err)
	}

	if blogFound == nil {
		t.Error("want a record, got nil")
	}
}
func testBlogsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Blogs(tx).Bind(blog); err != nil {
		t.Error(err)
	}
}

func testBlogsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Blogs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBlogsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogOne := &Blog{}
	blogTwo := &Blog{}
	if err = randomize.Struct(seed, blogOne, blogDBTypes, false, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}
	if err = randomize.Struct(seed, blogTwo, blogDBTypes, false, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = blogTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Blogs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBlogsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	blogOne := &Blog{}
	blogTwo := &Blog{}
	if err = randomize.Struct(seed, blogOne, blogDBTypes, false, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}
	if err = randomize.Struct(seed, blogTwo, blogDBTypes, false, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = blogTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Blogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testBlogsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Blogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlogsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx, blogColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Blogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlogToManyBlogPosts(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Blog
	var b, c BlogPost

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, blogPostDBTypes, false, blogPostColumnsWithDefault...)
	randomize.Struct(seed, &c, blogPostDBTypes, false, blogPostColumnsWithDefault...)

	b.BlogID = a.ID
	c.BlogID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	blogPost, err := a.BlogPosts(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range blogPost {
		if v.BlogID == b.BlogID {
			bFound = true
		}
		if v.BlogID == c.BlogID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BlogSlice{&a}
	if err = a.L.LoadBlogPosts(tx, false, (*[]*Blog)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BlogPosts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.BlogPosts = nil
	if err = a.L.LoadBlogPosts(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BlogPosts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", blogPost)
	}
}

func testBlogToManyAddOpBlogPosts(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Blog
	var b, c, d, e BlogPost

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blogDBTypes, false, strmangle.SetComplement(blogPrimaryKeyColumns, blogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BlogPost{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, blogPostDBTypes, false, strmangle.SetComplement(blogPostPrimaryKeyColumns, blogPostColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*BlogPost{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBlogPosts(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BlogID {
			t.Error("foreign key was wrong value", a.ID, first.BlogID)
		}
		if a.ID != second.BlogID {
			t.Error("foreign key was wrong value", a.ID, second.BlogID)
		}

		if first.R.Blog != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Blog != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.BlogPosts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.BlogPosts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.BlogPosts(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBlogsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = blog.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testBlogsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BlogSlice{blog}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testBlogsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Blogs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	blogDBTypes = map[string]string{`ID`: `bigint`, `Name`: `character varying`, `URL`: `character varying`}
	_           = bytes.MinRead
)

func testBlogsUpdate(t *testing.T) {
	t.Parallel()

	if len(blogColumns) == len(blogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Blogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	if err = blog.Update(tx); err != nil {
		t.Error(err)
	}
}

func testBlogsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(blogColumns) == len(blogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	blog := &Blog{}
	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Blogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, blog, blogDBTypes, true, blogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(blogColumns, blogPrimaryKeyColumns) {
		fields = blogColumns
	} else {
		fields = strmangle.SetComplement(
			blogColumns,
			blogPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(blog))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := BlogSlice{blog}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testBlogsUpsert(t *testing.T) {
	t.Parallel()

	if len(blogColumns) == len(blogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	blog := Blog{}
	if err = randomize.Struct(seed, &blog, blogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blog.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Blog: %s", err)
	}

	count, err := Blogs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &blog, blogDBTypes, false, blogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	if err = blog.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Blog: %s", err)
	}

	count, err = Blogs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
