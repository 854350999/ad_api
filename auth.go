package ad_api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//)
//

func Test() {

}

//func (c *Credential) GetGrantUrl() (string, error) {
//	hostUrl := c.getHostUrl()
//	fmt.Println("aaa")
//	return fmt.Sprintf("%s?client_id=%s&scope=%s&response_type=%s&redirect_uri=%s", hostUrl, c.ClientId, c.Scope, c.ResponseType, c.RedirectUri), nil
//}
//
//func (c *Credential) GetRefreshToken() (res response.Token, err error) {
//	//data := map[string]string{
//	//	"grant_type":    "authorization_code",
//	//	"code":          code,
//	//	"redirect_uri":  c.RedirectUri,
//	//	"client_id":     c.ClientId,
//	//	"client_secret": c.ClientSecret,
//	//}
//	//dataStr, err := json.Marshal(data)
//	//if err != nil {
//	//	return res, errors.New("failed to convert json")
//	//}
//	//tokenUrl := c.getTokenUrl()
//	//req, err := http.Post(tokenUrl)
//	return res, nil
//}
//
func (c *Client) GetNewAccessToken() (res TokenResponse, err error) {
	data := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": c.RefreshToken,
		"client_id":     c.ClientId,
		"client_secret": c.ClientSecret,
	}
	dataStr, err := json.Marshal(&data)
	if err != nil {
		return res, err
	}
	url := tokenUrl
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(dataStr))
	if err != nil {
		return res, err
	}
	resBody, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(resBody, &res); err != nil {
		return res, err
	}
	return res, nil
}
