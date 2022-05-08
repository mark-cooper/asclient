package asclient

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type PayLoadRequest struct {
	Method  string
	Path    string
	Payload string
	Params  QueryParams
}

type QueryParams struct {
	AllIds             bool     `url:"all_ids,omitempty"`
	Identifier         []string `url:"identifier,omitempty,brackets"`
	IdSet              string   `url:"id_set,omitempty"`
	IncludeDAOS        string   `url:"include_daos,omitempty"`
	IncludeUnpublished bool     `url:"include_unpublished,omitempty"`
	ModifiedSince      string   `url:"modified_since,omitempty"`
	NumberedCS         bool     `url:"numbered_cs,omitempty"`
	Page               int      `url:"page,omitempty"`
	Password           string   `url:"password,omitempty"`
	PrintPDF           bool     `url:"print_pdf,omitempty"`
	Query              string   `url:"q,omitempty"`
}

func (client *APIClient) BuildUrl(path string) (string, error) {
	rawUrl := client.CFG.URL + "/" + path
	u, err := url.Parse(rawUrl)

	if err != nil {
		return fmt.Sprintf("Unable to parse request URL: %s", rawUrl), err
	}
	return u.String(), nil
}

func (client *APIClient) CheckResponse(resp *resty.Response, err error) (*resty.Response, error) {
	if err != nil {
		return resp, err
	}

	if resp.StatusCode() != 200 {
		return resp, errors.New(string(resp.Body()))
	}

	return resp, nil
}

func (client *APIClient) Delete(path string) (*resty.Response, error) {
	url, _ := client.BuildUrl(path)

	resp, err := client.API.R().
		SetHeaders(client.Headers).
		Delete(url)

	return client.CheckResponse(resp, err)
}

func (client *APIClient) Get(path string, params QueryParams) (*resty.Response, error) {
	url, _ := client.BuildUrl(path)
	q, _ := query.Values(params)

	resp, err := client.API.R().
		SetHeaders(client.Headers).
		SetQueryString(q.Encode()).
		Get(url)

	return client.CheckResponse(resp, err)
}

func (client *APIClient) Post(path string, payload string, params QueryParams) (*resty.Response, error) {
	resp, err := client.RequestWithPayload(PayLoadRequest{
		Method:  "POST",
		Path:    path,
		Payload: payload,
		Params:  params,
	})

	return client.CheckResponse(resp, err)
}

func (client *APIClient) Put(path string, payload string, params QueryParams) (*resty.Response, error) {
	resp, err := client.RequestWithPayload(PayLoadRequest{
		Method:  "PUT",
		Path:    path,
		Payload: payload,
		Params:  params,
	})

	return client.CheckResponse(resp, err)
}

func (client *APIClient) RequestWithPayload(request PayLoadRequest) (*resty.Response, error) {
	url, _ := client.BuildUrl(request.Path)
	q, _ := query.Values(request.Params)

	resp, err := client.API.R().
		SetHeaders(client.Headers).
		SetContentLength(true).
		SetBody(request.Payload).
		SetQueryString(q.Encode()).
		Execute(request.Method, url)

	return resp, err
}
