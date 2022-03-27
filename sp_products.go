package ad_api

func (c *Client) GetSpCampaigns(r ...RequestOption) (res []byte, err error) {
	return c.PrepareRequest("GET", "/v2/sp/campaigns", "", r...)
}
