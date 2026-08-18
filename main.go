package main

import (
	"fmt"
	"net/http"
)

const address = ":8080"

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	fmt.Printf("server started at http://localhost%s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("server stopped", err)
	}
}
