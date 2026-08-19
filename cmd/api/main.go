package main

import (
	"fmt"
	"net/http"
	"os"
	"ticketswap-go/internal/database"
	"ticketswap-go/internal/events"
)

func main() {
	address := os.Getenv("SERVER_PORT")
	if address == "" {
		fmt.Println("SEVER_PORT env var is not set")
		return
	}

	db, err := database.OpenPostgres(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("failed to open Postgres:", err)
		return
	}
	defer db.Close()

	/*
	 * ---
	 * SETUP MUX
	 * ---
	 */

	mux := http.NewServeMux()

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
