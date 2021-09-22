package dingtalk

import "net/http"

// 根据userid获取用户详情  https://developers.dingtalk.com/document/app/query-user-details
func (c *Client) GetUserInfo(params *GetUserInfoParams) (*GetUserInfoRes, *http.Response, error) {
	req, err := c.NewRequest(http.MethodPost, "/topapi/v2/user/get", nil, params)
	if err != nil {
		return nil, nil, err
	}

	var data *GetUserInfoRes
	resp, err := c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}

// 获取部门用户详情 https://developers.dingtalk.com/document/app/queries-the-complete-information-of-a-department-user
func (c *Client) GetUserInfoWithDepartment(params *GetUserInfoWithDepartmentParams) (*GetUserInfoWithDepartmentRes, *http.Response, error) {
	req, err := c.NewRequest(http.MethodPost, "/topapi/v2/user/list", nil, params)
	if err != nil {
		return nil, nil, err
	}

	var data *GetUserInfoWithDepartmentRes
	resp, err := c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

// 获取部门用户userid列表 https://developers.dingtalk.com/document/app/query-the-list-of-department-userids
func (c *Client) GetUserIdWithDepartment(params *GetUserIdWithDepartmentParams) (*GetUserIdWithDepartmentRes, *http.Response, error) {
	req, err := c.NewRequest(http.MethodPost, "/topapi/user/listid", nil, params)
	if err != nil {
		return nil, nil, err
	}

	var data *GetUserIdWithDepartmentRes
	resp, err := c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}
