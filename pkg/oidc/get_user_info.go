package oidc

import (
	"fmt"
	"io"
	"net/http"
)

func GetUserInfo(configuration OpenIDConfiguration, token Token) error {
	req, err := http.NewRequest("GET", configuration.UserinfoEndpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
