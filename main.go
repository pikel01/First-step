package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Request struct {
	Text string `json:"text"`
}

type SuccessResponse struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/api/check", checkHandler)

	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{
			Error: "method not allowed",
		})
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{
			Error: "invalid json",
		})
		return
	}

	text := strings.TrimSpace(req.Text)

	if text != "hello" {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{
			Error: "only hello is allowed",
		})
		return
	}

	writeJSON(w, http.StatusOK, SuccessResponse{
		Result: "hello world",
	})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
