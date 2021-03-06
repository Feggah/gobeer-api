package beer_test

import (
	"database/sql"
	"testing"

	"github.com/feggah/gobeer-api/core/beer"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	_ "github.com/mattn/go-sqlite3"
)

const (
	errSetup  = "Error setting up the database:"
	errSearch = "Error searching in database:"
)

func cleanDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("delete from beer")
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func setupDB() (*beer.Service, error) {
	heineken := &beer.Beer{
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}

	becks := &beer.Beer{
		Name:  "Beck's",
		Type:  beer.TypeLager,
		Style: beer.StylePilsner,
	}

	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	if err != nil {
		return nil, err
	}
	// We won't defer here so we can run all tests without needing
	// to call sql.Open in every test case.

	err = cleanDB(db)
	if err != nil {
		return nil, err
	}

	service := beer.NewService(db)
	err = service.Store(heineken)
	if err != nil {
		return nil, err
	}

	err = service.Store(becks)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func TestGetAll(t *testing.T) {
	service, err := setupDB()
	if err != nil {
		t.Fatalf("%s %s", errSetup, err.Error())
	}

	beers, err := service.GetAll()
	if err != nil {
		t.Fatalf("%s %s", errSearch, err.Error())
	}
	if len(beers) != 2 {
		t.Fatalf("Error getting all beers. Expected length:%d, received: %d", 2, len(beers))
	}
}

func TestGet(t *testing.T) {
	heineken := &beer.Beer{
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	service, err := setupDB()
	if err != nil {
		t.Fatalf("%s %s", errSetup, err.Error())
	}

	beers, err := service.GetAll()
	if err != nil {
		t.Fatalf("%s %s", errSearch, err.Error())
	}
	var id int
	for _, i := range beers {
		if cmp.Equal(i, heineken, cmpopts.IgnoreFields(beer.Beer{}, "ID")) {
			id = i.ID
		}
	}

	b, err := service.Get(id)
	if err != nil {
		t.Fatalf("%s %s", errSearch, err.Error())
	}

	if diff := cmp.Diff(b, heineken, cmpopts.IgnoreFields(beer.Beer{}, "ID")); diff != "" {
		t.Errorf("TestGet: -want, +got:\n%s", diff)
	}
}

func TestStore(t *testing.T) {
	b := &beer.Beer{
		ID:    1,
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}

	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	if err != nil {
		t.Fatalf("Error connecting to the database: %s", err.Error())
	}
	defer db.Close()

	err = cleanDB(db)
	if err != nil {
		t.Fatalf("Error cleaning the database: %s", err.Error())
	}

	service := beer.NewService(db)
	err = service.Store(b)
	if err != nil {
		t.Fatalf("Error saving in database: %s", err.Error())
	}
}

func TestUpdate(t *testing.T) {
	heineken := &beer.Beer{
		ID:    1,
		Name:  "Heineken Lager Beer",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	service, err := setupDB()
	if err != nil {
		t.Fatalf("%s %s", errSetup, err.Error())
	}

	err = service.Update(heineken)
	if err != nil {
		t.Errorf("Error updating the database: %s", err.Error())
	}
}

func TestRemove(t *testing.T) {
	service, err := setupDB()
	if err != nil {
		t.Fatalf("%s %s", errSetup, err.Error())
	}

	err = service.Remove(1)
	if err != nil {
		t.Errorf("Error removing the database: %s", err.Error())
	}
}
