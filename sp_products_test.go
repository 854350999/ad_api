package ad_api

import "testing"

func TestClient_GetSpCampaigns(t *testing.T) {
	var cli = Client{}
	cli.AutoAccessToken = true

	cli.GetSpCampaigns(WithQueries(map[string]string{"count": "10"}))
}
