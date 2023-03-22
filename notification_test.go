package dingtalk

import (
	"testing"
)

func TestSendWorkerNotification(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	msg := TextMessage{
		Msgtype: "text",
		Text: Text{
			Content: "测试",
		},
	}

	params := &SendWorkerNotificationParams{
		UseridList: "",
		AgentID:    "",
		MSG:        msg,
	}

	data, _, err := client.SendWorkerNotification(params)
	if err != nil {
		t.Error(err)
	}

	t.Log(data)
}

func TestGetWorkerNotificationResult(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	params := &GetWorkerNotificationResultParams{
		TaskID:  0,
		AgentID: 0,
	}

	data, _, err := client.GetWorkerNotificationResult(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

type ActionCard struct {
	Msgtype    string           `json:"msgtype"`
	ActionCard ActionCardDetail `json:"action_card"`
}

type ActionCardDetail struct {
	Title       string `json:"title"`
	Markdown    string `json:"markdown"`
	SingleTitle string `json:"single_title"`
	SingleUrl   string `json:"single_url"`
}

func TestSendActionCardWorkerNotification(t *testing.T) {
	client, err := NewDingtalkClientWithParams("", "")
	if err != nil {
		t.Error(err)
	}

	message := ActionCard{
		Msgtype: "action_card",
		ActionCard: ActionCardDetail{
			Title:       "",
			Markdown:    "Markdown",
			SingleTitle: "点击进入工单系统",
			SingleUrl:   "",
		},
	}

	params := &SendWorkerNotificationParams{
		UseridList: "",
		AgentID:    "",
		MSG:        message,
	}

	data, _, err := client.SendWorkerNotification(params)
	if err != nil {
		t.Error(err)
	}

	t.Log(data)
}
