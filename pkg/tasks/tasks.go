package tasks

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var TaskServerUrl string = "https://tasks.internal.aleemhaji.com"

type TaskConfig struct {
	Endpoint string
	Headers  map[string]string
	Content  string
	Retries  int
	Delay    int
}

type TaskClientish interface {
	PostTask(TaskConfig) error
}

type TaskClient struct {
	ServerUrl string
}

func MakeTaskConfig(endpoint string) TaskConfig {
	return TaskConfig{
		Endpoint: endpoint,
		Headers:  map[string]string{},
		Content:  "",
		Retries:  3,
		Delay:    0,
	}
}

func MakeJsonTaskConfig(endpoint string, entity interface{}) TaskConfig {
	tc := MakeTaskConfig(endpoint)
	tc.Headers["Content-Type"] = "application/json"
	payload, err := json.Marshal(entity)
	if err != nil {
		panic(err)
	}

	tc.Content = string(payload)
	return tc
}

func NewDefaultTaskClient() *TaskClient {
	return NewTaskClient(TaskServerUrl)
}

func NewTaskClient(endpoint string) *TaskClient {
	return &TaskClient{
		ServerUrl: endpoint,
	}
}

func (t *TaskClient) PostTask(config TaskConfig) error {
	content, err := json.Marshal(config)
	if err != nil {
		return err
	}

	_, err = http.Post(t.ServerUrl, "application/json", bytes.NewBuffer(content))
	return err
}
