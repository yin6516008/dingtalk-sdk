package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

// NewDingtalkClient 创建一个钉钉客户端
func NewDingtalkClient(auth AuthInterface) (*Client, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	if token.Errcode != 0 {
		return nil, fmt.Errorf(token.Errmsg)
	}

	return &Client{
		accessToken:        token.AccessToken,
		accessTokenExpired: time.Now().Unix() + token.ExpiresIn,
		auth:               auth,
		baseUrl: &url.URL{
			Scheme: "https",
			Host:   "oapi.dingtalk.com",
		},
		http: &http.Client{},
	}, nil
}

// 创建一个Request
func (c *Client) NewRequest(method, path string, queryParams interface{}, bodyParams interface{}) (*http.Request, error) {

	if c.accessTokenExpired-time.Now().Unix() < 60 {
		token, err := c.auth.GetToken()
		if err != nil {
			return nil, err
		}

		if token.Errcode != 0 {
			return nil, fmt.Errorf(token.Errmsg)
		}

		c.accessToken = token.AccessToken
		c.accessTokenExpired = time.Now().Unix() + token.ExpiresIn
	}

	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}
	c.baseUrl.Path = unescaped

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	reqHeaders.Set("Content-Type", "application/json")

	var body []byte
	if bodyParams != nil {
		body, err = json.Marshal(bodyParams)
		if err != nil {
			return nil, err
		}
	}

	q, err := query.Values(queryParams)
	if err != nil {
		return nil, err
	}
	q.Set("access_token", c.accessToken)
	c.baseUrl.RawQuery = q.Encode()

	req, err := http.NewRequest(method, c.baseUrl.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

// 发送请求
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return resp, err
}
