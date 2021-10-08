package dingtalk

import "testing"

func TestListRole(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	params := &ListRoleParams{
		Size:   20,
		Offset: 0,
	}
	data, _, err := client.ListRole(params)
	if err != nil {
		t.Error(err)
	}

	t.Log(data)
}
