package oidc

import (
	"context"
	"log"
	"net/http"
	"time"
)

func GetAuthorizationCode() (string, error) {
	channel := make(chan string, 1)

	// configure server
	server := configureServer(&channel)

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
	err := shutdownServer(&server)
	if err != nil {
		return "", err
	}

	return authorizationCode, nil
}

func configureServer(channel *chan string) http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleCallbackFactory(channel))
	return http.Server{Addr: ":30423", Handler: mux}
}

func handleCallbackFactory(channel *chan string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You can now close this tab."))
		code := r.URL.Query().Get("code")
		*channel <- code
	}
}

func shutdownServer(server *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}
