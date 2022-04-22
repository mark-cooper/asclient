package asclient

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

func (client *ASpaceAPIClient) BuildUrl(path string) (string, error) {
	rawUrl := client.CFG.URL + "/" + path
	u, err := url.Parse(rawUrl)

	if err != nil {
		return fmt.Sprintf("Unable to parse request URL: %s", rawUrl), err
	}
	return u.String(), nil
}

func (client *ASpaceAPIClient) Post(path string, payload string, params map[string]string) (*resty.Response, error) {
	url, _ := client.BuildUrl(path)

	resp, err := client.API.R().
		SetHeaders(client.Headers).
		SetBody(payload).
		SetQueryParams(params).
		Post(url)

	return resp, err
}
