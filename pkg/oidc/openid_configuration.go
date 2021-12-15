package oidc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type OpenIDConfiguration struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
}

func GetOpenIDConfiguration() (OpenIDConfiguration, error) {

	url := "https://id.twitch.tv/oauth2/.well-known/openid-configuration"

	resp, err := http.Get(url)
	if err != nil {
		return OpenIDConfiguration{}, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return OpenIDConfiguration{}, readErr
	}

	openid_configuration := OpenIDConfiguration{}
	jsonErr := json.Unmarshal(body, &openid_configuration)
	if jsonErr != nil {
		return OpenIDConfiguration{}, jsonErr
	}

	return openid_configuration, nil
}
