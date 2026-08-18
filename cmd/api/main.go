package main

import (
	"fmt"
	"net/http"
	"ticketswap-go/internal/events"
)

const address = ":8080"

func main() {
	mux := http.NewServeMux()

	// GET Health
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("health endpoint called")
		fmt.Fprintln(w, "ok")
	})

	// Events
	eventHandler := events.Handler{}
	eventHandler.RegisterRoutes(mux)

	/*
	 * ---
	 * START THE SERVER
	 * ---
	 */

	fmt.Printf("server started at http://localhost%s\n", address)

	err := http.ListenAndServe(address, mux)
	if err != nil {
		fmt.Println("server stopped", err)
	}
}
