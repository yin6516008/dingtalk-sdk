package dingtalk

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	//client, err := NewDingtalkClientWithEnv()
	//if err != nil {
	//	t.Error(err)
	//}
	client, err := NewDingtalkClientWithParams("", "")
	if err != nil {
		t.Error(err)
	}
	params := &GetUserInfoParams{
		Userid: "",
	}
	data, _, err := client.GetUserInfo(params)
	if err != nil {
		t.Error(err)
	}
	marshal, err := json.Marshal(data.Result)
	if err != nil {
		return
	}
	t.Log(string(marshal))
	//	t.Log(data)
	//var org []string
	//	if len(data.Result.DeptIDList)==1 {
	//		org =DiGui(client,data.Result.DeptIDList[0],org)
	//		fmt.Println("1",org)
	//		Reverse(&org)
	//
	//	}
	//
	//	fmt.Println("2",	strings.Join(org,"-")+"｜"+data.Result.Title)
}

func DiGui(client *Client, depId int64, org []string) []string {

	params := &GetDepartmentInfoParams{
		DeptID: depId,
	}

	data, _, err := client.GetDepartmentInfo(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("qqq", data.Result.ParentID)
	org = append(org, data.Result.Name)
	fmt.Println(org)
	if data.Result.ParentID == 1 {
		return org
	}
	return DiGui(client, data.Result.ParentID, org)

}

func Reverse(arr *[]string) {
	var temp string
	length := len(*arr)
	for i := 0; i < length/2; i++ {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[length-1-i]
		(*arr)[length-1-i] = temp
	}
}
func TestGetUserInfoWithDepartment(t *testing.T) {
	//client, err := NewDingtalkClientWithEnv()
	client, err := NewDingtalkClientWithParams("", "")
	if err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err)
	}

	params := &GetUserInfoWithDepartmentParams{
		DeptID: 10086,
		Cursor: 0,
		Size:   100,
	}

	data, _, err := client.GetUserInfoWithDepartment(params)
	if err != nil {
		t.Error(err)
	}
	marshal, err := json.Marshal(data.Result)
	if err != nil {
		return
	}
	t.Log(string(marshal))
}

func TestGetUserIdWithDepartment(t *testing.T) {
	//client, err := NewDingtalkClientWithEnv()
	//if err != nil {
	//	t.Error(err)
	//}
	client, err := NewDingtalkClientWithParams("", "")
	if err != nil {
		t.Error(err)
	}

	params := &GetUserIdWithDepartmentParams{
		DeptID: 12,
	}
	data, _, err := client.GetUserIdWithDepartment(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(data.Result.UseridList)
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
