package main

import (
	"fmt"
	"os"

	"github.com/AmjedAgabani/qt-rup/pkg/oidc"
)

func main() {

	command := os.Args[1]

	switch command {
	case "help":
		help()
	case "login":
		openIdConfiguration := oidc.GetOpenIDConfiguration()
		oidc.OpenBrowser(openIdConfiguration)
		oidc.GetAuthorizationCode()
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
