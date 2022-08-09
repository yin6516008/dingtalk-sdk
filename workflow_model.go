package dingtalk

type StartWorkflowParams struct {
	AgentId             int                         `json:"agent_id,omitempty"`
	ProcessCode         string                      `json:"process_code"`
	OriginatorUserId    string                      `json:"originator_user_id"`
	DeptId              int                         `json:"dept_id"`
	Approvers           string                      `json:"approvers,omitempty"`
	ApproversV2         []ProcessInstanceApproverVo `json:"approvers_v2,omitempty"`
	CcList              string                      `json:"cc_list,omitempty"`
	CcPosition          string                      `json:"cc_position,omitempty"`
	FormComponentValues []FormComponentValueVo      `json:"form_component_values"`
}

type ProcessInstanceApproverVo struct {
	TaskActionType string   `json:"task_action_type,omitempty"`
	UserIds        []string `json:"user_ids,omitempty"`
}

type FormComponentValueVo struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	ExtValue string `json:"ext_value,omitempty"`
}

type StartWorkflowRes struct {
	BaseResponse
	ProcessInstanceID string `json:"process_instance_id"`
}
