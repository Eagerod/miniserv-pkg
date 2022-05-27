package tasks_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/Eagerod/miniserv-pkg/pkg/tasks"
)

func ExpectTasksPosted(t *testing.T, tcs []tasks.TaskConfig) (*httptest.Server, func() int) {
	remainingCalls := func() int {
		return len(tcs)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.NotEqual(t, 0, len(tcs))

		assert.Equal(t, "/", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var tc tasks.TaskConfig
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		err = json.Unmarshal(body, &tc)
		assert.NoError(t, err)

		assert.Equal(t, tcs[0], tc)
		tcs = tcs[1:]
	}))

	return server, remainingCalls
}
