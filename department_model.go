package dingtalk

type GetDepartmentInfoParams struct {
	Language string `json:"language"`
	DeptID   int64  `json:"dept_id"`
}

type GetDepartmentInfoRes struct {
	BaseResponse
	Result GetDepartmentInfoResult
}

type GetDepartmentInfoResult struct {
	DeptPermits         []int64       `json:"dept_permits"`
	OuterPermitUsers    []string      `json:"outer_permit_users"`
	OuterDept           bool          `json:"outer_dept"`
	GroupContainSubDept bool          `json:"group_contain_sub_dept"`
	DeptGroupChatID     string        `json:"dept_group_chat_id"`
	OrgDeptOwner        string        `json:"org_dept_owner"`
	AutoAddUser         bool          `json:"auto_add_user"`
	ParentID            int64         `json:"parent_id"`
	HideDept            bool          `json:"hide_dept"`
	Name                string        `json:"name"`
	OuterPermitDepts    []int64       `json:"outer_permit_depts"`
	UserPermits         []interface{} `json:"user_permits"`
	DeptID              int64         `json:"dept_id"`
	CreateDeptGroup     bool          `json:"create_dept_group"`
	Order               int64         `json:"order"`
}

type ListDepartmentParams struct {
	Language string `json:"language"`
	DeptID   int64  `json:"dept_id"`
}

type ListDepartmentRes struct {
	BaseResponse
	Result []ListDepartmentResult
}

type ListDepartmentResult struct {
	AutoAddUser     bool   `json:"auto_add_user"`
	CreateDeptGroup bool   `json:"create_dept_group"`
	DeptID          int64  `json:"dept_id"`
	Name            string `json:"name"`
	ParentID        int64  `json:"parent_id"`
}

// 获取子部门ID列表
type ListSubDepartmentIdParams struct {
	DeptID int64 `json:"dept_id"`
}

type ListSubDepartmentIdRes struct {
	BaseResponse
	Result ListSubDepartmentIdResult
}

type ListSubDepartmentIdResult struct {
	DeptIDList []int64 `json:"dept_id_list"`
}
