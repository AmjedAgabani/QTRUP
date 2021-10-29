package oidc

import (
	"fmt"
	"log"
	"net/http"
)

func handleCallback(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You can now close this tab."))
	code := r.URL.Query().Get("code")
	fmt.Println(code)
}

func GetAuthorizationCode() string {
	channel := make(chan string, 1)

	// channel <- "SOMEHOW CODE GOES THROUGH HERE"

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleCallback)

	go func() {
		server := &http.Server{Addr: ":30423", Handler: mux}
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	authorizationCode := <-channel

	fmt.Println(authorizationCode)

	return ""
}
