package main

import (
	"net/http"
	"os"
	"time"
)

func getPublicKeyHandler(publicKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(publicKey))
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	}
}
func main() {
	publicKey := os.Getenv("PUBLIC_KEY")
	if publicKey == "" {
		panic("PUBLIC_KEY environment variable is not set")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", getPublicKeyHandler(publicKey))

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
