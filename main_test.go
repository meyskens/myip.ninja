package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetIP(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, handleRequest(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "192.0.2.1", rec.Body.String())
	}
}

func TestGetIPJSON(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/?format=json", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, handleRequest(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"ip\":\"192.0.2.1\"}\n", rec.Body.String())
	}
}

func TestGetJSONP(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/?format=jsonp&callback=cb", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, handleRequest(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "cb({\"ip\":\"192.0.2.1\"}\n);", rec.Body.String())
	}
}

func TestGetXML(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/?format=xml", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, handleRequest(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<IP>192.0.2.1</IP>", rec.Body.String())
	}
}
