package oidc

import (
	"fmt"
	"log"
	"net/http"
)

func handleCallback(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You can now close this tab."))
	code := r.URL.Query().Get("code")
	fmt.Print(code)
}

func GetAuthorizationCode() string {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleCallback)
	err := http.ListenAndServe(":30423", mux)
	if err != nil {
		log.Fatal(err)
	}
	return ""
}
