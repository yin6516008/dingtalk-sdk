package dingtalk

import (
	"encoding/json"
	"testing"
)

func TestStartWorkflow(t *testing.T) {
	client, err := NewDingtalkClientWithParams("", "")
	if err != nil {
		t.Error(err)
	}

	params := &StartWorkflowParams{
		ProcessCode:      "",
		OriginatorUserId: "",
		DeptId:           0,
		Approvers:        "",
		ApproversV2: []ProcessInstanceApproverVo{
			{
				TaskActionType: "",
				UserIds:        []string{"", ""},
			},
		},
		CcList:     "",
		CcPosition: "",
		FormComponentValues: []FormComponentValueVo{
			{
				Name:  "",
				Value: "",
			},
			{
				Name:  "",
				Value: "",
			},
			{
				Name:  "",
				Value: "",
			},
			{
				Name:  "",
				Value: "",
			},
			{
				Name:  "",
				Value: "",
			},
			{
				Name:  "",
				Value: "",
			},
		},
	}

	data, _, err := client.StartWorkflow(params)
	if err != nil {
		t.Error(err)
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	t.Log(string(marshal))
}
