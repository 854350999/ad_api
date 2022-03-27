package ad_api

func (c *Client) GetPortfolios(r ...RequestOption) (res []byte, err error) {
	return c.PrepareRequest("GET", "/v2/portfolios", "", r...)
}
