package dingtalk

import (
	"net/http"
)

// 通过免登码获取用户信息 https://open.dingtalk.com/document/orgapp-server/obtain-the-userid-of-a-user-by-using-the-log-free
func (c *Client) GetUserInfoByFreeCode(params *GetUserInfoByFreeCodeParams) (data *GetUserInfoByFreeCodeRes, resp *http.Response, err error) {
	req, err := c.NewRequest(http.MethodPost, "/topapi/v2/user/getuserinfo", nil, params)
	if err != nil {
		return nil, nil, err
	}

	resp, err = c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}
