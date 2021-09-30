package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	router = chi.NewRouter()

)

func TestMain(m *testing.M){
	RegisterHandlers(router)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestAddBlog(t *testing.T) {

	ts := httptest.NewServer(router)
	defer ts.Close()

	req,_:= http.NewRequest(http.MethodGet, ts.URL +"/",nil)

	res, _ :=http.DefaultClient.Do(req)

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected http.StatusOK but got res.StatusC")
	}
}

