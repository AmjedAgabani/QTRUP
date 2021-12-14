package oidc

import (
	"encoding/json"
	"log"
	"os"
)

func SaveToken(token Token) {

	bytes, err := json.Marshal(token)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("token.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write(bytes)
}

// save token as txt file
// convert token from struct to string
// write string into txt file

// load token
// find existing txt file
// copy the string inside
// convert back to struct
