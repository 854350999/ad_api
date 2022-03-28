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
