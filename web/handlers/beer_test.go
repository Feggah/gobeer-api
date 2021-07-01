package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/feggah/gobeer-api/core/beer"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

const (
	endpoint = "/v1/beer"
)

type BeerServiceMock struct{}

func (t BeerServiceMock) GetAll() ([]*beer.Beer, error) {
	b1 := &beer.Beer{
		ID:    10,
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	b2 := &beer.Beer{
		ID:    20,
		Name:  "Skol",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	return []*beer.Beer{b1, b2}, nil
}

func (t BeerServiceMock) Get(ID int) (*beer.Beer, error) {
	b1 := &beer.Beer{
		ID:    ID,
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	return b1, nil
}
func (t BeerServiceMock) Store(b *beer.Beer) error {
	return nil
}

func (t BeerServiceMock) Update(b *beer.Beer) error {
	return nil
}

func (t BeerServiceMock) Remove(ID int) error {
	return nil
}

func TestGetAllBeer(t *testing.T) {
	handler := getAllBeer(&BeerServiceMock{})

	r := mux.NewRouter()
	r.Handle(endpoint, handler)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	assert.Nil(t, err)
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var result []*beer.Beer
	err = json.NewDecoder(rr.Body).Decode(&result)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
}

func TestGetBeer(t *testing.T) {
	var id int = 1

	handler := getBeer(&BeerServiceMock{})
	r := mux.NewRouter()
	r.Handle(fmt.Sprintf("%s/{id}", endpoint), handler)
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%d", endpoint, id),
		nil,
	)
	assert.Nil(t, err)
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var result *beer.Beer
	err = json.NewDecoder(rr.Body).Decode(&result)
	assert.Nil(t, err)
	assert.Equal(t, id, result.ID)
	assert.Equal(t, "Heineken", result.Name)
}

func TestStoreBeer(t *testing.T) {
	b := beer.Beer{
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	beerJson, err := json.Marshal(b)
	assert.Nil(t, err)

	handler := storeBeer(&BeerServiceMock{})
	r := mux.NewRouter()
	r.Handle(endpoint, handler)
	req, err := http.NewRequest(
		http.MethodPost,
		endpoint,
		strings.NewReader(string(beerJson)),
	)
	assert.Nil(t, err)
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestUpdateBeer(t *testing.T) {
	id := 1
	b := beer.Beer{
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	beerJson, err := json.Marshal(b)
	assert.Nil(t, err)

	handler := updateBeer(&BeerServiceMock{})
	r := mux.NewRouter()
	r.Handle(fmt.Sprintf("%s/{id}", endpoint), handler)
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/%d", endpoint, id),
		strings.NewReader(string(beerJson)),
	)
	assert.Nil(t, err)
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestRemoveBeer(t *testing.T) {
	id := 1

	handler := removeBeer(&BeerServiceMock{})
	r := mux.NewRouter()
	r.Handle(fmt.Sprintf("%s/{id}", endpoint), handler)
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/%d", endpoint, id),
		nil,
	)
	assert.Nil(t, err)
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestValidateData(t *testing.T) {
	b := &beer.Beer{}
	v := validateData(*b)

	assert.Equal(t, v, false)
}

func TestMakeBeerHandlers(t *testing.T) {
	var endpoints []string = []string{
		endpoint,
		fmt.Sprintf("%s/1", endpoint),
	}
	service := &BeerServiceMock{}
	r := mux.NewRouter()
	n := negroni.New()

	MakeBeerHandlers(r, n, service)

	for _, i := range endpoints {
		req, err := http.NewRequest(http.MethodOptions, i, nil)
		assert.Nil(t, err)
		req.Header.Set("Accept", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	}
}
