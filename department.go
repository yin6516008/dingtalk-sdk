package dingtalk

import "net/http"

// 获取部门详情  https://developers.dingtalk.com/document/app/query-department-details0-v2
func (c *Client) GetDepartmentInfo(params *GetDepartmentInfoParams) (*GetDepartmentInfoRes, *http.Response, error) {
	req, err := c.NewRequest(http.MethodPost, "/topapi/v2/department/get", nil, params)
	if err != nil {
		return nil, nil, err
	}

	var data *GetDepartmentInfoRes
	resp, err := c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}
	return data, resp, err

}

// 获取部门列表 https://developers.dingtalk.com/document/app/obtain-the-department-list-v2
func (c *Client) ListDepartment(params *ListDepartmentParams) (*ListDepartmentRes, *http.Response, error) {
	req, err := c.NewRequest(http.MethodPost, "/topapi/v2/department/listsub", nil, params)
	if err != nil {
		return nil, nil, err
	}

	var data *ListDepartmentRes
	resp, err := c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}
	return data, resp, err
}

// 获取子部门ID列表 https://developers.dingtalk.com/document/app/obtain-a-sub-department-id-list-v2
func (c *Client) ListSubDepartmentId(params *ListSubDepartmentIdParams) (*ListSubDepartmentIdRes, *http.Response, error) {
	req, err := c.NewRequest(http.MethodPost, "/topapi/v2/department/listsubid", nil, params)
	if err != nil {
		return nil, nil, err
	}

	var data *ListSubDepartmentIdRes
	resp, err := c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}
