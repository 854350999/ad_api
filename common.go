package ad_api

func (c *Client) GetPortfolios(r ...RequestOption) (res []byte, err error) {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/v2/portfolios",
	}
	return c.DoRequest(preReq, r...)
}
