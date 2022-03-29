package ad_api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"unsafe"
)

type Credentials struct {
	ClientId     string
	ClientSecret string
	Region       string
	RefreshToken string
}

type Client struct {
	Credentials
	AccessToken     string
	ProfileId       int64
	AutoAccessToken bool
}

func NewClient(cre Credentials) *Client {
	return &Client{
		Credentials: cre,
	}
}

type PrepareRequest struct {
	Method      string
	UriPath     string
	ContentType string
}

type Request struct {
	Queries map[string]string
	Body    io.Reader
}

type RequestOption func(*Request)

func WithQueries(queries map[string]string) RequestOption {
	return func(request *Request) {
		request.Queries = queries
	}
}

func WithBody(body io.Reader) RequestOption {
	return func(request *Request) {
		request.Body = body
	}
}

func NewRequest(options ...RequestOption) *Request {
	request := &Request{
		Queries: nil,
		Body:    nil,
	}
	for _, o := range options {
		o(request)
	}
	return request
}

func (c *Client) getEndpoint() (string, error) {
	endpoint := endpoints[c.Region]
	if endpoint == "" {
		return "", errors.New("a invalid region")
	}
	return endpoint, nil
}

func (c *Client) DoRequest(preReq PrepareRequest, options ...RequestOption) *Response {
	res := &Response{}
	if c.AutoAccessToken == true {
		tokenRes, err := c.GetNewAccessToken()
		if err != nil {
			res.Error = err
			return res
		}
		c.AccessToken = tokenRes.AccessToken
	}
	endpoint, err := c.getEndpoint()
	if err != nil {
		res.Error = err
		return res
	}
	var queries []string
	r := NewRequest(options...)
	if r.Queries != nil {
		for k, v := range r.Queries {
			query := fmt.Sprintf("%s=%s", k, v)
			queries = append(queries, query)
		}
	}
	queryPath := strings.Join(queries, "&")
	if queryPath != "" {
		queryPath = "?" + queryPath
	}
	url := endpoint + preReq.UriPath + queryPath
	req, err := http.NewRequest(preReq.Method, url, r.Body)
	if preReq.ContentType == "" {
		preReq.ContentType = "application/json"
	}
	req.Header.Set("Content-Type", preReq.ContentType)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	req.Header.Set("Amazon-Advertising-API-ClientId", c.ClientId)
	req.Header.Set("Amazon-Advertising-API-Scope", strconv.FormatInt(c.ProfileId, 10))
	client := http.Client{}
	resp, err := client.Do(req.WithContext(context.TODO()))
	if err != nil {
		res.Error = err
		return res
	}
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Error = err
		return res
	}
	res.StatusCode = resp.StatusCode
	res.Payload = respByte
	res.CheckResponse()
	return res
}

func (res *Response) CheckResponse() {
	if res.StatusCode >= 400 {
		err := *(*string)(unsafe.Pointer(&res.Payload))
		res.Error = errors.New(err)
	}
}
