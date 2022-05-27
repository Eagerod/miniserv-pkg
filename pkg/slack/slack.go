package slack

import (
	"github.com/Eagerod/miniserv-pkg/pkg/tasks"
)

var SlackBotUrl string = "https://slackbot.internal.aleemhaji.com/message"
var AlertingChannel string = "CKE1AKEAV"
var DefaultChannel string = "CHB1UECGJ"

type SlackTaskClient tasks.TaskClient


func (t *SlackTaskClient) SendMessage(message string) error {
	return t.SendMessageChannel(message, DefaultChannel)
}

func (t *SlackTaskClient) SendMessageAlert(message string) error {
	return t.SendMessageChannel(message, AlertingChannel)
}

func (t *SlackTaskClient) SendMessageChannel(message, channel string) error {
	taskConfig := tasks.MakeTaskConfig(SlackBotUrl)
	taskConfig.Headers["Content-Type"] = "text/plain"
	taskConfig.Headers["X-SLACK-CHANNEL-ID"] = channel
	taskConfig.Content = message
	
	return t.PostTask(taskConfig)
}

func (t *SlackTaskClient) PostTask(config tasks.TaskConfig) error {
	return (*tasks.TaskClient)(t).PostTask(config)
}
