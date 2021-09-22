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
		Userid: "1223575014880951",
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
		DeptID: 59009774,
		Size:   10,
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
		DeptID: 332373921,
	}
	data, _, err := client.GetUserIdWithDepartment(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

func TestName(t *testing.T) {
	client, err := NewDingtalkClient(&CorpInternalAuth{AppKey: "dingljjybah4q1noub1o", AppSecret: "zi9Fsr6THZS1uG0Inn6JD2-x3MKf6w8dY5m5nmc0euKjzbGaeX0JekBcXZF4uylr"})
	if err != nil {
		t.Error(err)
	}

	params := &GetUserInfoParams{
		Userid: "1223575014880951",
	}
	data, _, err := client.GetUserInfo(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}
