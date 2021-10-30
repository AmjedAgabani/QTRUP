package oidc

import (
	"context"
	"log"
	"net/http"
	"time"
)

func GetAuthorizationCode() string {
	channel := make(chan string, 1)

	// configure server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// get code
		w.Write([]byte("You can now close this tab."))
		code := r.URL.Query().Get("code")
		// send code to main thread
		channel <- code
	})
	server := &http.Server{Addr: ":30423", Handler: mux}

	// start server
	go func() {
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// wait for authorization code
	authorizationCode := <-channel

	// shutdown server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return authorizationCode
}
