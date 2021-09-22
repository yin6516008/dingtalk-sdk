package dingtalk

import (
	"fmt"
	"os"
	"testing"
)

func TestGetToken(t *testing.T) {
	var appKey string
	if appKey = os.Getenv("DINGTALK_APP_KEY"); appKey == "" {
		t.Error("Not DINGTALK_APP_KEY environment variable")
	}

	var appSecret string
	if appSecret = os.Getenv("DINGTALK_APP_SECRET"); appSecret == "" {
		t.Error("Not DINGTALK_APP_SECRET environment variable")
	}

	auth := CorpInternalAuth{
		AppKey:    appKey,
		AppSecret: appSecret,
	}

	data, err := auth.GetToken()
	if err != nil {
		t.Error(err)
	}

	if data.Errcode != 0 {
		t.Error(err)
	}
	fmt.Println(data)
}
