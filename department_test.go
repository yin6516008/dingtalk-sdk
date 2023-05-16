package dingtalk

import (
	"encoding/json"
	"testing"
)

func TestGetDepartmentInfo(t *testing.T) {
	//client, err := NewDingtalkClientWithEnv()
	//if err != nil {
	//	t.Error(err)
	//}
	client, err := NewDingtalkClientWithParams("", "")
	if err != nil {
		t.Error(err)
	}

	params := &GetDepartmentInfoParams{
		DeptID: 1,
	}

	data, _, err := client.GetDepartmentInfo(params)
	if err != nil {
		t.Error(err)
	}
	marshal, err := json.Marshal(data.Result)
	if err != nil {
		return
	}
	t.Log(string(marshal))

	t.Log(data)

}

func TestListDepartment(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}
	params := ListDepartmentParams{
		DeptID: 11,
	}
	data, _, err := client.ListDepartment(&params)
	if err != nil {
		t.Error(err)
	}
	t.Log(data.Result[0].DeptID)
}

func TestListSubDepartmentId(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	params := &ListSubDepartmentIdParams{
		DeptID: 11,
	}
	data, _, err := client.ListSubDepartmentId(params)
	if err != nil {
		t.Error(err)
	}

	t.Log(data)
}

func TestGetAllParentDepartmentListOfUser(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	params := &GetAllParentDepartmentListOfUserParams{
		Userid: "",
	}
	data, _, err := client.GetAllParentDepartmentListOfUser(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}
