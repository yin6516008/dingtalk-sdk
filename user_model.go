package dingtalk

// 根据userid获取用户详情
type GetUserInfoParams struct {
	Language string `json:"language"`
	Userid   string `json:"userid"`
}

type GetUserInfoRes struct {
	BaseResponse
	Result User `json:"result"`
}

type User struct {
	Extension     string          `json:"extension"`
	Unionid       string          `json:"unionid"`
	Boss          bool            `json:"boss"`
	UnionEmpEXT   UnionEmpEXT     `json:"unionEmpExt"`
	RoleList      []RoleList      `json:"role_list"`
	Admin         bool            `json:"admin"`
	Remark        string          `json:"remark"`
	Title         string          `json:"title"`
	HiredDate     int64           `json:"hired_date"`
	Userid        string          `json:"userid"`
	WorkPlace     string          `json:"work_place"`
	DeptOrderList []DeptOrderList `json:"dept_order_list"`
	RealAuthed    bool            `json:"real_authed"`
	DeptIDList    []int64         `json:"dept_id_list"`
	JobNumber     string          `json:"job_number"`
	Email         string          `json:"email"`
	OrgMail       string          `json:"org_email"`
	LeaderInDept  []LeaderInDept  `json:"leader_in_dept"`
	ManagerUserid string          `json:"manager_userid"`
	Mobile        string          `json:"mobile"`
	Active        bool            `json:"active"`
	Telephone     string          `json:"telephone"`
	Avatar        string          `json:"avatar"`
	HideMobile    bool            `json:"hide_mobile"`
	Senior        bool            `json:"senior"`
	Name          string          `json:"name"`
	StateCode     string          `json:"state_code"`
}

type DeptOrderList struct {
	DeptID int64   `json:"dept_id"`
	Order  float64 `json:"order"`
}

type LeaderInDept struct {
	DeptID int64 `json:"dept_id"`
	Leader bool  `json:"leader"`
}

type RoleList struct {
	GroupName string `json:"group_name"`
	ID        int64  `json:"id"`
	Name      string `json:"name"`
}

type UnionEmpEXT struct {
	CorpID          string            `json:"corpId"`
	Userid          string            `json:"userid"`
	UnionEmpMapList []UnionEmpMapList `json:"unionEmpMapList"`
}

type UnionEmpMapList struct {
	CorpID string `json:"corpId"`
	Userid string `json:"userid"`
}

// 获取部门用户详情
type GetUserInfoWithDepartmentParams struct {
	Cursor             int64  `json:"cursor"`
	ContainAccessLimit bool   `json:"contain_access_limit"`
	Size               int64  `json:"size"`
	OrderField         string `json:"order_field"`
	Language           string `json:"language"`
	DeptID             int64  `json:"dept_id"`
}

type GetUserInfoWithDepartmentRes struct {
	BaseResponse
	Result GetUserInfoWithDepartmentResult
}

type GetUserInfoWithDepartmentResult struct {
	PaginationInfo
	List []GetUserInfoWithDepartmentList `json:"list"`
}

type GetUserInfoWithDepartmentList struct {
	DeptOrder        int64   `json:"dept_order"`
	Leader           bool    `json:"leader"`
	Extension        string  `json:"extension"`
	Unionid          string  `json:"unionid"`
	Boss             bool    `json:"boss"`
	ExclusiveAccount bool    `json:"exclusive_account"`
	Mobile           string  `json:"mobile"`
	Active           bool    `json:"active"`
	Admin            bool    `json:"admin"`
	Telephone        string  `json:"telephone"`
	Remark           string  `json:"remark"`
	Avatar           string  `json:"avatar"`
	HideMobile       bool    `json:"hide_mobile"`
	Title            string  `json:"title"`
	HiredDate        int64   `json:"hired_date"`
	Userid           string  `json:"userid"`
	WorkPlace        string  `json:"work_place"`
	OrgEmail         string  `json:"org_email"`
	Name             string  `json:"name"`
	DeptIDList       []int64 `json:"dept_id_list"`
	JobNumber        string  `json:"job_number"`
	StateCode        string  `json:"state_code"`
	Email            string  `json:"email"`
}

// 获取部门用户userid列表
type GetUserIdWithDepartmentParams struct {
	DeptID int64 `json:"dept_id"`
}

type GetUserIdWithDepartmentRes struct {
	BaseResponse
	Result GetUserIdWithDepartmentResult
}

type GetUserIdWithDepartmentResult struct {
	UseridList []string `json:"userid_list"`
}
