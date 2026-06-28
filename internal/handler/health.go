package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp := HealthResponse{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("encode response error: %v", err)
	}
}
