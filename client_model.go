package dingtalk

import (
	"net/http"
	"net/url"
)

type Client struct {
	AccessToken        string
	accessTokenExpired int64
	auth               AuthInterface
	baseUrl            *url.URL
	http               *http.Client
}

type BaseResponse struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	RequestID string `json:"request_id"`
}

type PaginationInfo struct {
	NextCursor string `json:"next_cursor"`
	HasMore    bool   `json:"has_more"`
}
