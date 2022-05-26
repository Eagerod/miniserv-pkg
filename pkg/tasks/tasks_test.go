package tasks

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

func TestPostTask(t *testing.T) {
	calls := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var tc TaskConfig
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		err = json.Unmarshal(body, &tc)
		assert.NoError(t, err)

		assert.Equal(t, "https://nothing.internal.aleemhaji.com", tc.Endpoint)
		assert.Equal(t, map[string]string{}, tc.Headers)
		assert.Equal(t, "", tc.Content)
		assert.Equal(t, 3, tc.Retries)
		assert.Equal(t, 0, tc.Delay)

		calls += 1
	}))
	defer server.Close()

	config := MakeTaskConfig("https://nothing.internal.aleemhaji.com")
	NewTaskClient(server.URL).PostTask(config)

	assert.Equal(t, 1, calls)
}

func TestJsonPostTask(t *testing.T) {
	calls := 0

	type T struct {
		A string
		B int
	}

	payload := T{
		"Yeah",
		123,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var tc TaskConfig
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		err = json.Unmarshal(body, &tc)
		assert.NoError(t, err)

		content := tc.Content
		var receivedPayload T
		err = json.Unmarshal([]byte(content), &receivedPayload)
		assert.NoError(t, err)

		expectedHeaders := map[string]string{
			"Content-Type": "application/json",
		}

		assert.Equal(t, "https://nothing.internal.aleemhaji.com", tc.Endpoint)
		assert.Equal(t, expectedHeaders, tc.Headers)
		assert.Equal(t, payload, receivedPayload)
		assert.Equal(t, 3, tc.Retries)
		assert.Equal(t, 0, tc.Delay)

		calls += 1
	}))
	defer server.Close()

	config := MakeJsonTaskConfig("https://nothing.internal.aleemhaji.com", payload)
	NewTaskClient(server.URL).PostTask(config)

	assert.Equal(t, 1, calls)
}
