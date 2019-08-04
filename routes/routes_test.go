package routes

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	bandname "github.com/davidbanham/bandname_go"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/", serveIndex).Methods("GET")

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestSubmitHandler(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest("GET", "/submit", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/submit", serveSubmit).Methods("GET")

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestJokeCreateHandler(t *testing.T) {
	t.Parallel()

	form := url.Values{
		"name":      {bandname.Bandname()},
		"body":      {bandname.Bandname()},
		"submitter": {bandname.Bandname()},
		"score":     {"10"},
	}
	req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Path: "/jokes"},
		Form:   form,
	}

	rr := httptest.NewRecorder()
	jokeCreateHandler(rr, req)

	assert.Equal(t, http.StatusFound, rr.Code)
}

func TestJokeCreateHandlerLowScore(t *testing.T) {
	t.Parallel()

	form := url.Values{
		"name":      {bandname.Bandname()},
		"body":      {bandname.Bandname()},
		"submitter": {bandname.Bandname()},
		"score":     {"1"},
	}
	req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Path: "/jokes"},
		Form:   form,
	}

	rr := httptest.NewRecorder()
	jokeCreateHandler(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestJokeCreateHandlerNegativeScore(t *testing.T) {
	t.Parallel()

	form := url.Values{
		"name":      {bandname.Bandname()},
		"body":      {bandname.Bandname()},
		"submitter": {bandname.Bandname()},
		"score":     {"-1"},
	}
	req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Path: "/jokes"},
		Form:   form,
	}

	rr := httptest.NewRecorder()
	jokeCreateHandler(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestJokeCreateHandlerMassiveScore(t *testing.T) {
	t.Parallel()

	form := url.Values{
		"name":      {bandname.Bandname()},
		"body":      {bandname.Bandname()},
		"submitter": {bandname.Bandname()},
		"score":     {"999999999"},
	}
	req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Path: "/jokes"},
		Form:   form,
	}

	rr := httptest.NewRecorder()
	jokeCreateHandler(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestJokeCreateHandlerHighScore(t *testing.T) {
	t.Parallel()

	form := url.Values{
		"name":      {bandname.Bandname()},
		"body":      {bandname.Bandname()},
		"submitter": {bandname.Bandname()},
		"score":     {"100"},
	}
	req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Path: "/jokes"},
		Form:   form,
	}

	rr := httptest.NewRecorder()
	jokeCreateHandler(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
