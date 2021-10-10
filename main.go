package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	command := os.Args[1]

	switch command {
	case "help":
		help()
	case "login":
		logIn()
	default:
		fmt.Println("Something went wrong")
	}
}

func help() {
	fmt.Println(`Quick Twitch Random User Picker.

Usage:

    qt-rup <command> [arguments]

The commands are:
	
    help      shows all commands & their uses
    stats     see your statistics about your current subscribers
    export    export a beautified csv list of current subscribers
    select    select a random subscriber / gifter
    login     log into your twitch Account (opens twitch oauth page) 
    logout    logout of your twitch account without closing the application
    exit      exit the app (you will automatically be logged out)`)
}

type openid_configuration struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
}

func logIn() {

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

	openid_configuration := openid_configuration{}
	jsonErr := json.Unmarshal(body, &openid_configuration)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(openid_configuration)
}
