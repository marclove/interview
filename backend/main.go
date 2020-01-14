package main

import (
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/api", helloHandler)
	http.ListenAndServe("localhost:4000", nil)
}

type Payload struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	data := Payload{Message: "Hello another"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
