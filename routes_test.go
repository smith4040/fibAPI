package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fibNum uint = 0
var r = setupRouter(&fibNum)

func TestNegativeNumberHandling(t *testing.T) {
	req, _ := http.NewRequest("GET", "/previous", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, "text/plain; charset=utf-8", w.HeaderMap.Get("Content-Type"))
	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "API does not return negative numbers.", w.Body.String())
}

func TestServeCurrentNumber(t *testing.T) {
	req, _ := http.NewRequest("GET", "/current", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, "text/plain; charset=utf-8", w.HeaderMap.Get("Content-Type"))
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "0", w.Body.String())
}

func TestServeNextNumber(t *testing.T) {
	req, _ := http.NewRequest("GET", "/next", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, "text/plain; charset=utf-8", w.HeaderMap.Get("Content-Type"))
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "1", w.Body.String())
}

func TestServePreviousNumber(t *testing.T) {
	req, _ := http.NewRequest("GET", "/previous", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, "text/plain; charset=utf-8", w.HeaderMap.Get("Content-Type"))
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "0", w.Body.String())
}
