package main

import (
	"fmt"
	"net/http"
	"ticketswap-go/internal/database"
	"ticketswap-go/internal/events"
)

const address = ":8080"
const postgresURI = "postgres://ticketswap:localdev@localhost:5432/ticketswap?sslmode=disable"

func main() {
	mux := http.NewServeMux()

	db, err := database.OpenPostgres(postgresURI)
	if err != nil {
		fmt.Println("failed to open Postgres:", err)
		return
	}
	defer db.Close()

	// GET Health
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("health endpoint called")
		fmt.Fprintln(w, "ok")
	})

	// Events
	eventHandler := events.Handler{DB: db}
	eventHandler.RegisterRoutes(mux)

	/*
	 * ---
	 * START THE SERVER
	 * ---
	 */

	fmt.Printf("server started at http://localhost%s\n", address)

	err = http.ListenAndServe(address, mux)
	if err != nil {
		fmt.Println("server stopped", err)
	}
}
