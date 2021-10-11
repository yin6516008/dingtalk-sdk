package dingtalk

type GetUserInfoByFreeCodeParams struct {
	Code string `json:"code" url:"code"`
}

type GetUserInfoByFreeCodeRes struct {
	BaseResponse
	AssociatedUnionid string `json:"associated_unionid"`
	Unionid           string `json:"unionid"`
	DeviceID          string `json:"device_id"`
	SysLevel          int64  `json:"sys_level"`
	Name              string `json:"name"`
	Sys               bool   `json:"sys"`
	Userid            string `json:"userid"`
}
