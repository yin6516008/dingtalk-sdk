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
	//client, err := NewDingtalkClientWithEnv()
	//if err != nil {
	//	t.Error(err)
	//}
	client, err := NewDingtalkClientWithParams("", "")
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
	marshal, err := json.Marshal(data.Result)
	if err != nil {
		return
	}
	t.Log(string(marshal))

	t.Log(data)
}

func TestListParentByDept(t *testing.T) {

	client, err := NewDingtalkClientWithParams("", "")
	if err != nil {
		t.Error(err)
	}

	params := &ListParentByDeptParams{DeptID: 666}

	data, _, err := client.ListParentByDept(params)
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

var list []int64

func TestListAllSubDept(t *testing.T) {
	list = []int64{}
	client, err := NewDingtalkClientWithParams("", "")
	if err != nil {
		t.Error(err)
	}
	//var list []int64
	ListAllSubDept(client, 1)
}

func ListAllSubDept(client *Client, depId int64) {
	data, err := fetchData(client, depId)
	if err != nil {
		return
	}

	list = append(list, data...)
	if len(data) != 0 {
		for _, v := range data {
			ListAllSubDept(client, v)
		}
	}

	return
}

func fetchData(client *Client, depId int64) ([]int64, error) {
	params := &ListSubDepartmentIdParams{DeptID: depId}
	info, _, err := client.ListSubDepartmentId(params)
	if err != nil {
		return nil, err
	}
	return info.Result.DeptIDList, nil
}
