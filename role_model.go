package dingtalk

// 获取角色列表
type ListRoleParams struct {
	Size   int64 `json:"size"`
	Offset int64 `json:"offset"`
}

type ListRoleRes struct {
	BaseResponse
	Result ListRoleResult
}

type ListRoleResult struct {
	HasMore bool `json:"hasMore"`
	List    []ListRoleList
}

type ListRoleList struct {
	GroupID int64          `json:"groupId"`
	Name    string         `json:"name"`
	Roles   []ListRoleRole `json:"roles"`
}

type ListRoleRole struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
