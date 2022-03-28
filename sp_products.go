package ad_api

import (
	"fmt"
	"strconv"
)

func (c *Client) CreateSpCampaigns(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: "/v2/sp/campaigns",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) UpdateSpCampaigns(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "PUT",
		UriPath: "/v2/sp/campaigns",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSpCampaigns(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/v2/sp/campaigns",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSpCampaignById(campaignId int) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/v2/sp/campaigns/%s", strconv.Itoa(campaignId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) GetSpCampaignsExtended(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/v2/sp/campaigns/extended",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSpCampaignExtendedById(campaignId int) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/v2/sp/campaigns/extended/%s", strconv.Itoa(campaignId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) DeleteSpCampaignById(campaignId int) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "DELETE",
		UriPath: fmt.Sprintf("/v2/sp/campaigns/extended/%s", strconv.Itoa(campaignId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) CreateSpAdGroups(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: "/v2/sp/adGroups",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) UpdateSpAdGroups(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "PUT",
		UriPath: "/v2/sp/adGroups",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSpAdGroups(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/v2/sp/adGroups",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSpAdGroupById(adGroupId int) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/v2/sp/adGroups/%s", strconv.Itoa(adGroupId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) GetSpAdGroupsExtended(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/v2/sp/adGroups/extended",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSpAdGroupExtendedById(adGroupId int) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/v2/sp/adGroups/extended/%s", strconv.Itoa(adGroupId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) DeleteSpAdGroupById(adGroupId int) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "DELETE",
		UriPath: fmt.Sprintf("/v2/sp/adGroups/%s", strconv.Itoa(adGroupId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) GetSpBidRecommendationByAdGroupId(adGroupId int) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/v2/sp/adGroups/%s/bidRecommendations", strconv.Itoa(adGroupId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) GetSpBidRecommendationByKeywordId(keywordId int) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/v2/sp/keywords/%s/bidRecommendations", strconv.Itoa(keywordId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) GetSpKeywordsBidRecommendations(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: "/v2/sp/keywords/bidRecommendations",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSpTargetsBidRecommendations(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: "/v2/sp/targets/bidRecommendations",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) CreateSpKeywords(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: "/v2/sp/keywords",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) UpdateSpKeywords(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "PUT",
		UriPath: "/v2/sp/keywords",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSpKeywords(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "PUT",
		UriPath: "/v2/sp/keywords",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSpKeywordById(keywordId int) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/v2/sp/keywords/%s", strconv.Itoa(keywordId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) GetSpKeywordsExtended(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "PUT",
		UriPath: "/v2/sp/keywords/extended",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSpKeywordExtendedById(keywordId int) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/v2/sp/keywords/extended/%s", strconv.Itoa(keywordId)),
	}
	return c.DoRequest(preReq)
}
