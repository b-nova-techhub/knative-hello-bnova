package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("knative-hello-bnova: received a request")
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", message)
}

func main() {
	log.Print("knative-hello-bnova: starting server...")

	http.HandleFunc("/", handler)

	port := "8080"

	log.Printf("knative-hello-bnova: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
