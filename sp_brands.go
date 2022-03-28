package ad_api

import (
	"fmt"
	"strconv"
)

func (c *Client) CreateSbCampaigns(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: "/sb/campaigns",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) UpdateSbCampaigns(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "PUT",
		UriPath: "/sb/campaigns",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSbCampaigns(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/sb/campaigns",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSbCampaignById(campaignId int) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/sb/campaigns/%s", strconv.Itoa(campaignId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) DeleteSbCampaignById(campaignId int) *Response {
	preReq := PrepareRequest{
		Method:  "DELETE",
		UriPath: fmt.Sprintf("/sb/campaigns/%s", strconv.Itoa(campaignId)),
	}
	return c.DoRequest(preReq)
}
