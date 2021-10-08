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
		UseridList: "1223575014880951",
		AgentID:    "1217926382",
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
		AgentID: 1217926382,
	}

	data, _, err := client.GetWorkerNotificationResult(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}
