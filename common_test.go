package ad_api

import "testing"

func TestClient_GetPortfolios(t *testing.T) {
	var cli = Client{}
	cli.AutoAccessToken = true
	cli.GetPortfolios()
}
