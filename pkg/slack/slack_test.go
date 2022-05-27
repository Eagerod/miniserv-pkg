package slack

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/Eagerod/miniserv-pkg/pkg/tasks"
	"github.com/Eagerod/miniserv-pkg/pkg/tasks/test"
)

func TestSendMessage(t *testing.T) {
	expectedCall := tasks.MakeTaskConfig("https://example.com/message")
	expectedCall.Content = "some message content"
	expectedCall.Headers["Content-Type"] = "text/plain"
	expectedCall.Headers[SlackChannelHttpHeaderName] = DefaultChannel

	server, rc := tasks_test.ExpectTasksPosted(t, []tasks.TaskConfig{expectedCall})
	defer server.Close()

	slackTaskClient := NewSlackTaskClient("https://example.com/message", server.URL)
	slackTaskClient.SendMessage("some message content")

	assert.Equal(t, 0, rc())
}
