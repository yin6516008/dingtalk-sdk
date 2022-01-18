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

type GetUserInfoByCodeQeury struct {
	AccessKey string `json:"accessKey" url:"accessKey"`
	Timestamp string `json:"timestamp" url:"timestamp"`
	Signature string `json:"signature" url:"signature"`
}

type GetUserInfoByCodeRes struct {
	BaseResponse
	Nick                 string `json:"nick"`
	Unionid              string `json:"unionid"`
	Openid               string `json:"openid"`
	MainOrgAuthHighLevel bool   `json:"main_org_auth_high_level"`
}
