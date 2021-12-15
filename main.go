package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AmjedAgabani/qt-rup/pkg/oidc"
)

func main() {

	command := os.Args[1]

	switch command {
	case "help":
		help()
	case "login":
		openIdConfiguration, err := oidc.GetOpenIDConfiguration()
		if err != nil {
			log.Fatal(err)
		}

		err = oidc.OpenBrowser(openIdConfiguration)
		if err != nil {
			log.Fatal(err)
		}

		authorizationCode, err := oidc.GetAuthorizationCode()
		if err != nil {
			log.Fatal(err)
		}

		token, err := oidc.PostToken(openIdConfiguration, authorizationCode)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(token)

		err = oidc.SaveToken(token)
		if err != nil {
			log.Fatal(err)
		}

	case "stats":
		token, err := oidc.LoadToken()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(token)
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
