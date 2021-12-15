package oidc

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func SaveToken(token Token) {

	bytes, err := json.Marshal(token)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("token.txt", bytes, 0600)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadToken() Token {
	bytes, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatal(err)
	}
	token := Token{}
	jsonErr := json.Unmarshal(bytes, &token)
	if err != nil {
		log.Fatal(jsonErr)
	}

	return token
}
