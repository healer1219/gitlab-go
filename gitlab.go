package gitlab_go

import (
	"net/url"
)

type GitlabClient struct {
	BaseUrl      *url.URL
	PrivateToken string
}

const (
	defaultUrl         = "https://gitlab.com"
	privateTokenHeader = "PRIVATE_TOKEN"
)

var gitlab GitlabClient

func WithAccessToken(targetUrl string, accessToken string) (*GitlabClient, error) {
	if gitlab != (GitlabClient{}) && accessToken == gitlab.PrivateToken && targetUrl == gitlab.BaseUrl.String() {
		return &gitlab, nil
	}
	return nil, nil
}
