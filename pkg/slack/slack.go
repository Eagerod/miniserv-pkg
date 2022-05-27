package slack

import (
	"github.com/Eagerod/miniserv-pkg/pkg/tasks"
)

var SlackBotUrl string = "https://slackbot.internal.aleemhaji.com/message"
var AlertingChannel string = "CKE1AKEAV"
var DefaultChannel string = "CHB1UECGJ"
var SlackChannelHttpHeaderName string = "X-Slack-Channel-Id"

type SlackTaskClient struct {
	*tasks.TaskClient

	SlackEndpoint string
}

func NewDefaultSlackTaskClient() *SlackTaskClient {
	return &SlackTaskClient{
		tasks.NewDefaultTaskClient(),
		SlackBotUrl,
	}
}

func NewSlackTaskClient(slackEndpoint, tasksEndpoint string) *SlackTaskClient {
	return &SlackTaskClient{
		TaskClient: tasks.NewTaskClient(tasksEndpoint),
		SlackEndpoint: slackEndpoint,
	}
}

func (t *SlackTaskClient) SendMessage(message string) error {
	return t.SendMessageChannel(message, DefaultChannel)
}

func (t *SlackTaskClient) SendMessageAlert(message string) error {
	return t.SendMessageChannel(message, AlertingChannel)
}

func (t *SlackTaskClient) SendMessageChannel(message, channel string) error {
	taskConfig := tasks.MakeTaskConfig(t.SlackEndpoint)
	taskConfig.Headers["Content-Type"] = "text/plain"
	taskConfig.Headers[SlackChannelHttpHeaderName] = channel
	taskConfig.Content = message

	return t.TaskClient.PostTask(taskConfig)
}
