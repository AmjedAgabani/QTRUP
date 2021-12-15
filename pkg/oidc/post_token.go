package oidc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Token struct {
	AccessToken  string   `json:"access_token"`
	ExpiresIn    uint64   `json:"expires_in"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
	TokenType    string   `json:"token_type"`
}

func PostToken(configuration OpenIDConfiguration, authorizationCode string) (Token, error) {

	token := Token{}

	data := url.Values{
		"client_id":     {"gx35ttsecbzblr2ksi9a7l2beqsx8c"},
		"client_secret": {"0g78pf7gg7t0fh7o7vjkrx07ndrnqq"},
		"code":          {authorizationCode},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {"http://localhost:30423"},
	}

	resp, err := http.PostForm(configuration.TokenEndpoint, data)
	if err != nil {
		return token, err
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return token, readErr
	}

	jsonErr := json.Unmarshal(body, &token)
	if jsonErr != nil {
		return token, jsonErr
	}

	return token, nil

}
