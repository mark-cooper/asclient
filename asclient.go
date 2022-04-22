package asclient

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/go-resty/resty/v2"
)

type ASpaceAPIClient struct {
	API     resty.Client
	CFG     ASpaceAPIConfig
	Headers map[string]string
}

type ASpaceAPIConfig struct {
	URL      string
	Username string
	Password string
}

type ASpaceAPISessionResponse struct {
	Token string `json:"session"`
}

func NewAPIClient(config ASpaceAPIConfig) ASpaceAPIClient {
	return ASpaceAPIClient{
		API: *resty.New(),
		CFG: config,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func (client *ASpaceAPIClient) BuildUrl(path string) (string, error) {
	rawUrl := client.CFG.URL + "/" + path
	u, err := url.Parse(rawUrl)

	if err != nil {
		return fmt.Sprintf("Unable to parse request URL: %s", rawUrl), err
	}
	return u.String(), nil
}

func (client *ASpaceAPIClient) Login() (string, error) {
	resp, _ := client.Post(
		filepath.Join("users", client.CFG.Username, "login"),
		"{}",
		map[string]string{
			"password": client.CFG.Password,
		},
	)

	session := ASpaceAPISessionResponse{}
	json.Unmarshal(resp.Body(), &session)

	client.Headers["X-ArchivesSpace-Session"] = session.Token
	return session.Token, nil
}

func (client *ASpaceAPIClient) Post(path string, payload string, params map[string]string) (*resty.Response, error) {
	url, _ := client.BuildUrl(path)

	resp, _ := client.API.R().
		SetHeaders(client.Headers).
		SetBody(payload).
		SetQueryParams(params).
		Post(url)

	return resp, nil
}
