package dingtalk

import "testing"

func TestGetUserInfoByFreeCode(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	params := &GetUserInfoByFreeCodeParams{
		Code: "faewfawef",
	}
	data, _, err := client.GetUserInfoByFreeCode(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}
