package ad_api

import (
	"fmt"
	"strconv"
)

func (c *Client) CreateSdCampaigns(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: "/sd/campaigns",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) UpdateSdCampaigns(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "PUT",
		UriPath: "/sd/campaigns",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSdCampaigns(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/sd/campaigns",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSdCampaignById(campaignId int) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/sd/campaigns/%s", strconv.Itoa(campaignId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) GetSdCampaignsExtended(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/sd/campaigns/extended",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSdCampaignExtendedById(campaignId int) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/sd/campaigns/extended/%s", strconv.Itoa(campaignId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) CreateSdAdGroups(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: "/sd/adGroups",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) UpdateSdAdGroups(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "PUT",
		UriPath: "/sd/adGroups",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSdAdGroups(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/sd/adGroups",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSdAdGroupById(adGroupId int) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/sd/adGroups/%s", strconv.Itoa(adGroupId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) GetSdAdGroupsExtended(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/sd/adGroups/extended",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSdAdGroupExtendedById(adGroupId int) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/sd/adGroups/extended/%s", strconv.Itoa(adGroupId)),
	}
	return c.DoRequest(preReq)
}
