package http_json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Return a status code, and message that could be returned to a client after
//   parsing out a JSON object from an HTTP request.
// Mostly to save boilerplate content in other handlers
func JsonUnmarshalBody(r *http.Request, entity interface{}) (int, error) {
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "application/json") {
		return http.StatusUnsupportedMediaType, fmt.Errorf("JSON payload required")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("failed to read body")
	}

	if len(body) == 0 {
		return http.StatusBadRequest, fmt.Errorf("must include a body")
	}

	if err := json.Unmarshal(body, entity); err != nil {
		return http.StatusBadRequest, fmt.Errorf("invalid JSON provided")
	}

	return http.StatusOK, nil
}
