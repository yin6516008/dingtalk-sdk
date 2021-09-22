package dingtalk

type GetTokenRes struct {
	BaseResponse
	AccessToken string `json:"access_token"`
	Errmsg      string `json:"errmsg"`
	ExpiresIn   int64  `json:"expires_in"`
}

type AuthInterface interface {
	GetToken() (*GetTokenRes, error)
}
