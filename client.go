package ad_api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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

type Request struct {
	Queries map[string]string
	Body    map[string]interface{}
}

func (c *Client) getEndpoint() (string, error) {
	endpoint := endpoints[c.Region]
	if endpoint == "" {
		return "", errors.New("a invalid region")
	}
	return endpoint, nil
}

func (c *Client) PrepareRequest(method string, uriPath string, contentType string, r Request) (res []byte, err error) {
	if c.AutoAccessToken == true {
		tokenRes, err := c.GetNewAccessToken()
		if err != nil {
			return res, err
		}
		c.AccessToken = tokenRes.AccessToken
	}
	endpoint, err := c.getEndpoint()
	if err != nil {
		return res, err
	}
	var queries []string
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
	url := endpoint + uriPath + queryPath
	bodyByte, err := json.Marshal(&r.Body)
	if err != nil {
		return res, err
	}
	if r.Body == nil {
		bodyByte = nil
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyByte))
	if contentType == "" {
		contentType = "application/json"
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	req.Header.Set("Amazon-Advertising-API-ClientId", c.ClientId)
	req.Header.Set("Amazon-Advertising-API-Scope", strconv.FormatInt(c.ProfileId, 10))
	client := http.Client{}
	resp, err := client.Do(req.WithContext(context.TODO()))
	if err != nil {
		return res, err
	}
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	return respByte, nil
}
