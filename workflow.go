package dingtalk

import "net/http"

func (c *Client) StartWorkflow(params *StartWorkflowParams) (data *StartWorkflowRes, resp *http.Response, err error) {
	req, err := c.NewRequest(http.MethodPost, "/topapi/processinstance/create", nil, params)
	if err != nil {
		return nil, nil, err
	}

	resp, err = c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}
	return data, resp, err
}
