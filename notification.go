package dingtalk

import "net/http"

func (c *Client) SendWorkerNotification(params *SendWorkerNotificationParams) (data *SendWorkerNotificationRes, resp *http.Response, err error) {
	req, err := c.NewRequest(http.MethodPost, "/topapi/message/corpconversation/asyncsend_v2", nil, params)
	if err != nil {
		return nil, nil, err
	}

	resp, err = c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}
	return data, resp, err
}

func (c *Client) GetWorkerNotificationResult(params *GetWorkerNotificationResultParams) (data *GetWorkerNotificationResultRes, resp *http.Response, err error) {
	req, err := c.NewRequest(http.MethodPost, "/topapi/message/corpconversation/getsendresult", nil, params)
	if err != nil {
		return nil, nil, err
	}

	resp, err = c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}
	return data, resp, err
}
