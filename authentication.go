package dingtalk

import "net/http"

func (c *Client) GetUserInfoByFreeCode(params *GetUserInfoByFreeCodeParams) (data *GetUserInfoByFreeCodeRes, resp *http.Response, err error) {
	req, err := c.NewRequest(http.MethodGet, "/topapi/v2/user/getuserinfo", params, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}
