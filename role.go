package dingtalk

import (
	"net/http"
)

func (c *Client) ListRole(params *ListRoleParams) (*ListRoleRes, *http.Response, error) {
	req, err := c.NewRequest(http.MethodPost, "/topapi/role/list", nil, params)
	if err != nil {
		return nil, nil, err
	}

	var data *ListRoleRes
	resp, err := c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}
	return data, resp, err
}
