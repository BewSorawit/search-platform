package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"search-platform/internal/milvus"
	"time"
)

type ReadyHandler struct {
	mc *milvus.Client
}

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func writeJSONError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Status: "not_ready",
		Error:  msg,
	}); err != nil {
		log.Printf("write json error failed: %v", err)
	}

}

func NewReadyHandler(mc *milvus.Client) *ReadyHandler {
	return &ReadyHandler{mc: mc}
}

func (h *ReadyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := h.mc.Ping(ctx); err != nil {
		writeJSONError(w, http.StatusServiceUnavailable, "milvus not ready")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{
		"status": "ready",
	}); err != nil {
		log.Printf("encode ready response error: %v", err)
	}

}
