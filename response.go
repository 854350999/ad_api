package ad_api

type Response struct {
	StatusCode int
	Payload    []byte
	Error      error
}

type TokenResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

type PortfoliosResponse struct {
	PortfolioId int    `json:"portfolioId"`
	Name        string `json:"name"`
	InBudget    bool   `json:"inBudget"`
	State       string `json:"state"`
}
