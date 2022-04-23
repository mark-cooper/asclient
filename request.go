package asclient

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type PayLoadRequest struct {
	Method  string
	Path    string
	Payload string
	Params  map[string]string
}

func (client *ASpaceAPIClient) BuildUrl(path string) (string, error) {
	rawUrl := client.CFG.URL + "/" + path
	u, err := url.Parse(rawUrl)

	if err != nil {
		return fmt.Sprintf("Unable to parse request URL: %s", rawUrl), err
	}
	return u.String(), nil
}

func (client *ASpaceAPIClient) Delete(path string) (*resty.Response, error) {
	url, _ := client.BuildUrl(path)

	resp, err := client.API.R().
		SetHeaders(client.Headers).
		Delete(url)

	return resp, err
}

func (client *ASpaceAPIClient) Get(path string, params map[string]string) (*resty.Response, error) {
	url, _ := client.BuildUrl(path)

	resp, err := client.API.R().
		SetHeaders(client.Headers).
		SetQueryParams(params).
		Get(url)

	return resp, err
}

func (client *ASpaceAPIClient) Post(path string, payload string, params map[string]string) (*resty.Response, error) {
	resp, err := client.RequestWithPayload(PayLoadRequest{
		Method:  "POST",
		Path:    path,
		Payload: payload,
		Params:  params,
	})

	return resp, err
}

func (client *ASpaceAPIClient) Put(path string, payload string, params map[string]string) (*resty.Response, error) {
	resp, err := client.RequestWithPayload(PayLoadRequest{
		Method:  "PUT",
		Path:    path,
		Payload: payload,
		Params:  params,
	})

	return resp, err
}

func (client *ASpaceAPIClient) RequestWithPayload(request PayLoadRequest) (*resty.Response, error) {
	url, _ := client.BuildUrl(request.Path)

	resp, err := client.API.R().
		SetHeaders(client.Headers).
		SetContentLength(true).
		SetBody(request.Payload).
		SetQueryParams(request.Params).
		Execute(request.Method, url)

	return resp, err
}
