package oidc

import (
	"encoding/json"
	"io/ioutil"
)

func SaveToken(token Token) error {

	bytes, err := json.Marshal(token)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("token.txt", bytes, 0600)
	if err != nil {
		return err
	}

	return nil
}

func LoadToken() (Token, error) {
	token := Token{}

	bytes, err := ioutil.ReadFile("token.txt")
	if err != nil {
		return token, err
	}

	jsonErr := json.Unmarshal(bytes, &token)
	if jsonErr != nil {
		return token, jsonErr
	}

	return token, nil
}
