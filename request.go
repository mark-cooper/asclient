package asclient

import (
	"encoding/json"
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

type QueryParams struct {
	AllIds             string `json:"all_ids,omitempty"`
	IdSet              string `json:"id_set,omitempty"`
	IncludeDAOS        string `json:"include_daos,omitempty"`
	IncludeUnpublished string `json:"include_unpublished,omitempty"`
	ModifiedSince      string `json:"modified_since,omitempty"`
	NumberedCS         string `json:"numbered_cs,omitempty"`
	Page               string `json:"page,omitempty"`
	Password           string `json:"password,omitempty"`
	PrintPDF           string `json:"print_pdf,omitempty"`
}

func (client *ASpaceAPIClient) BuildUrl(path string) (string, error) {
	rawUrl := client.CFG.URL + "/" + path
	u, err := url.Parse(rawUrl)

	if err != nil {
		return fmt.Sprintf("Unable to parse request URL: %s", rawUrl), err
	}
	return u.String(), nil
}

func (client *ASpaceAPIClient) ConvertParams(params QueryParams) (map[string]string, error) {
	var queryStringParams map[string]string
	data, _ := json.Marshal(params)
	json.Unmarshal(data, &queryStringParams)
	return queryStringParams, nil
}

func (client *ASpaceAPIClient) Delete(path string) (*resty.Response, error) {
	url, _ := client.BuildUrl(path)

	resp, err := client.API.R().
		SetHeaders(client.Headers).
		Delete(url)

	return resp, err
}

func (client *ASpaceAPIClient) Get(path string, params QueryParams) (*resty.Response, error) {
	url, _ := client.BuildUrl(path)
	queryStringParams, _ := client.ConvertParams(params)

	resp, err := client.API.R().
		SetHeaders(client.Headers).
		SetQueryParams(queryStringParams).
		Get(url)

	return resp, err
}

func (client *ASpaceAPIClient) Post(path string, payload string, params QueryParams) (*resty.Response, error) {
	queryStringParams, _ := client.ConvertParams(params)

	resp, err := client.RequestWithPayload(PayLoadRequest{
		Method:  "POST",
		Path:    path,
		Payload: payload,
		Params:  queryStringParams,
	})

	return resp, err
}

func (client *ASpaceAPIClient) Put(path string, payload string, params QueryParams) (*resty.Response, error) {
	queryStringParams, _ := client.ConvertParams(params)

	resp, err := client.RequestWithPayload(PayLoadRequest{
		Method:  "PUT",
		Path:    path,
		Payload: payload,
		Params:  queryStringParams,
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
