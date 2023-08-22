package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hostname := os.Getenv("VIRTUAL_HOST")
	if hostname == "" {
		fmt.Fprintf(os.Stdout, "VIRTUAL_HOST is empty! Shutting down...")
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Listening on :8080\n")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
		body := "<!DOCTYPE html><html><head><title>" + hostname + "</title></head><body><center><h1>Welcome to " + hostname + "!</h1></center></body></html>"
		fmt.Fprintf(w, body)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
