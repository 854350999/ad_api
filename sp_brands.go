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

func (c *Client) GetSbAdGroups(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/sb/adGroups",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSbAdGroupById(adGroupId int) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/sb/adGroups/%s", strconv.Itoa(adGroupId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) CreateSbKeywords(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: "/sb/keywords",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) UpdateSbKeywords(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "PUT",
		UriPath: "/sb/keywords",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSbKeywords(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/sb/keywords",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSbKeywordById(keywordId int) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/sb/keywords/%s", strconv.Itoa(keywordId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) DeleteSbKeywordById(keywordId int) *Response {
	preReq := PrepareRequest{
		Method:  "DELETE",
		UriPath: fmt.Sprintf("/sb/keywords/%s", strconv.Itoa(keywordId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) CreateSbTargets(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: "/sb/targets",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) UpdateSbTargets(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "PUT",
		UriPath: "/sb/targets",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSbTargets(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: "/sb/targets/list",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSbTargetById(targetId int) *Response {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: fmt.Sprintf("/sb/targets/%s", strconv.Itoa(targetId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) DeleteSbTargetById(targetId int) *Response {
	preReq := PrepareRequest{
		Method:  "DELETE",
		UriPath: fmt.Sprintf("/sb/targets/%s", strconv.Itoa(targetId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) GetSbBidRecommendations(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: "/sb/recommendations/bids",
	}
	return c.DoRequest(preReq, r...)
}
