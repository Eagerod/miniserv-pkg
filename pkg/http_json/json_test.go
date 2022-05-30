package http_json

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

type TestEntity struct {
	A string
	B int
}

func testHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var payload TestEntity
	status, err := JsonUnmarshalBody(r, &payload)
	if err != nil {
		w.WriteHeader(status)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func TestJsonUnmarshalBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testHandlerFunc))
	defer server.Close()

	payload := []byte(`{"A":"b","B":1}`)
	reader := bytes.NewReader(payload)
	r, err := http.Post(server.URL, "application/json", reader)
	assert.NoError(t, err)

	responsePayload, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.Equal(t, []byte("OK"), responsePayload)
}

func TestJsonUnmarshalBodyWrongContentType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testHandlerFunc))
	defer server.Close()

	payload := []byte(`{"A":"b","B":1}`)
	reader := bytes.NewReader(payload)
	r, err := http.Post(server.URL, "text/plain", reader)
	assert.NoError(t, err)

	responsePayload, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusUnsupportedMediaType, r.StatusCode)
	assert.Equal(t, []byte("JSON payload required"), responsePayload)
}

func TestJsonUnmarshalBodyEmptyBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testHandlerFunc))
	defer server.Close()

	payload := []byte{}
	reader := bytes.NewReader(payload)
	r, err := http.Post(server.URL, "application/json", reader)
	assert.NoError(t, err)

	responsePayload, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, r.StatusCode)
	assert.Equal(t, []byte("must include a body"), responsePayload)
}

func TestJsonUnmarshalBodyInvalidJson(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testHandlerFunc))
	defer server.Close()

	payload := []byte("I promise I'm JSON")
	reader := bytes.NewReader(payload)
	r, err := http.Post(server.URL, "application/json", reader)
	assert.NoError(t, err)

	responsePayload, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, r.StatusCode)
	assert.Equal(t, []byte("invalid JSON provided"), responsePayload)
}
