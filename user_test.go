package dingtalk

import (
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	params := &GetUserInfoParams{
		Userid: "XXX",
	}
	data, _, err := client.GetUserInfo(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

func TestGetUserInfoWithDepartment(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	params := &GetUserInfoWithDepartmentParams{
		DeptID: 5900,
		Size:   12,
	}

	data, _, err := client.GetUserInfoWithDepartment(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

func TestGetUserIdWithDepartment(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	params := &GetUserIdWithDepartmentParams{
		DeptID: 234455,
	}
	data, _, err := client.GetUserIdWithDepartment(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

func TestGetUseridByUnionid(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	data, _, err := client.GetUserIdByUnionid("xxx")
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

func TestGetUserinfoByCode(t *testing.T) {
	client, err := NewDingtalkClientWithParams("", "")
	if err != nil {
		t.Error(err)
	}

	data, _, err := client.GetUserInfoByCode("")
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}
