package dingtalk

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

func NewDingtalkClientWithEnv() (*Client, error) {
	var appKey string
	if appKey = os.Getenv("DINGTALK_APP_KEY"); appKey == "" {
		return nil, fmt.Errorf("not DINGTALK_APP_KEY environment variable")
	}

	var appSecret string
	if appSecret = os.Getenv("DINGTALK_APP_SECRET"); appSecret == "" {
		return nil, fmt.Errorf("not DINGTALK_APP_SECRET environment variable")
	}

	auth := &CorpInternalAuth{
		AppKey:    appKey,
		AppSecret: appSecret,
	}

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

func TestCorpClient(t *testing.T) {
	client, err := NewDingtalkClientWithEnv()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(client.accessToken)

}
