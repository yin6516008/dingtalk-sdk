package dingtalk

import (
	"testing"
)

func TestGetDepartmentInfo(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	params := &GetDepartmentInfoParams{
		DeptID: 59009774,
	}

	data, _, err := client.GetDepartmentInfo(params)
	if err != nil {
		t.Error(err)
	}

	t.Log(data)

}

func TestListDepartment(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}
	params := ListDepartmentParams{
		DeptID: 59009774,
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
		DeptID: 59009774,
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
		Userid: "1223575014880951",
	}
	data, _, err := client.GetAllParentDepartmentListOfUser(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}
