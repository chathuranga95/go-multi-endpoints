package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response structure for JSON format
type Response struct {
	Hello string `json:"hello"`
	Port  int    `json:"port"`
}

// Handler for the first service on port 9090
func handler9090(w http.ResponseWriter, r *http.Request) {
	response := Response{Hello: "world", Port: 9090}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler for the second service on port 9091
func handler9091(w http.ResponseWriter, r *http.Request) {
	response := Response{Hello: "world", Port: 9091}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Setting up the first service listener on port 9090
	http.HandleFunc("/service1", handler9090)
	go func() {
		if err := http.ListenAndServe(":9090", nil); err != nil {
			fmt.Printf("Error starting server on port 9090: %v\n", err)
		}
	}()

	// Setting up the second service listener on port 9091
	http.HandleFunc("/service2", handler9091)
	go func() {
		if err := http.ListenAndServe(":9091", nil); err != nil {
			fmt.Printf("Error starting server on port 9091: %v\n", err)
		}
	}()

	// Prevent the main function from exiting
	select {}
}
