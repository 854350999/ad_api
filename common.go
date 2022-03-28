package ad_api

func (c *Client) GetPortfolios(r ...RequestOption) *Response {
	preReq := PrepareRequest{
		Method:  "GET",
		UriPath: "/v2/portfolios",
	}
	return c.DoRequest(preReq, r...)
}
