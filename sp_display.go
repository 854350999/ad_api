package ad_api

import (
	"fmt"
	"strconv"
)

func (c *Client) CreateSdCampaigns(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "POST",
		UriPath: "/sd/campaigns",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) UpdateSdCampaigns(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "PUT",
		UriPath: "/sd/campaigns",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSdCampaigns(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/sd/campaigns",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSdCampaignById(campaignId int) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/sd/campaigns/%s", strconv.Itoa(campaignId)),
	}
	return c.DoRequest(preReq)
}

func (c *Client) GetSdCampaignsExtended(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/sd/campaigns/extended",
	}
	return c.DoRequest(preReq, r...)
}

func (c *Client) GetSdCampaignExtendedById(campaignId int) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: fmt.Sprintf("/sd/campaigns/extended/%s", strconv.Itoa(campaignId)),
	}
	return c.DoRequest(preReq)
}
