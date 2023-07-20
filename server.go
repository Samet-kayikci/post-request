package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Text struct {
	Str string `json:"text"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("data not post\nr.Method:", r.Method)
		return
	}
	var data Text
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "error.", http.StatusBadRequest)
		return
	}
	fmt.Println("data:", data.Str)

	response := "connected"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/connect", handler)
	fmt.Println("HTTP server port number :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error server:", err)
	}
}
