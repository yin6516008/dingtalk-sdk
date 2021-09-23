package dingtalk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 企业内部应用
type CorpInternalAuth struct {
	AppKey    string
	AppSecret string
}

func (c *CorpInternalAuth) GetToken() (*GetTokenRes, error) {
	query := url.Values{}
	query.Set("appkey", c.AppKey)
	query.Set("appsecret", c.AppSecret)

	url := url.URL{
		Scheme:   "https",
		Host:     "oapi.dingtalk.com",
		Path:     "gettoken",
		RawQuery: query.Encode(),
	}

	resp, err := http.Get(url.String())

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data *GetTokenRes
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if data.Errcode != 0 {
		return nil, fmt.Errorf("get token faild, code: %d, msg: %s ", data.Errcode, data.Errmsg)
	}

	return data, err
}
