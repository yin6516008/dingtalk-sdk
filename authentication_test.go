package dingtalk

import (
	"testing"
)

func TestGetUserInfoByFreeCode(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	params := &GetUserInfoByFreeCodeParams{
		Code: "f3167a9fbd0131cba8e45eeb6298f263",
	}
	data, _, err := client.GetUserInfoByFreeCode(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}
