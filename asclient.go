package asclient

import (
	"encoding/json"
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

func (client *ASpaceAPIClient) BuildUrl(path []string) string {
	return client.CFG.URL + "/" + filepath.Join(path...)
}

func (client *ASpaceAPIClient) Login() (string, error) {
	path := client.BuildUrl([]string{"users", client.CFG.Username, "login"})

	resp, _ := client.API.R().
		SetHeaders(client.Headers).
		SetQueryParams(map[string]string{
			"password": client.CFG.Password,
		}).
		Post(path)

	session := ASpaceAPISessionResponse{}
	json.Unmarshal(resp.Body(), &session)

	client.Headers["X-ArchivesSpace-Session"] = session.Token
	return session.Token, nil
}
