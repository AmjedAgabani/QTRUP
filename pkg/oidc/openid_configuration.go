package oidc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type OpenIDConfiguration struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
}

func GetOpenIDConfiguration() OpenIDConfiguration {

	url := "https://id.twitch.tv/oauth2/.well-known/openid-configuration"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	openid_configuration := OpenIDConfiguration{}
	jsonErr := json.Unmarshal(body, &openid_configuration)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}


	return openid_configuration
}
