package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"search-platform/internal/milvus"
	"time"
)

type IngestHandler struct {
	mc *milvus.Client
}

func NewIngestHandler(mc *milvus.Client) *IngestHandler {
	return &IngestHandler{mc: mc}
}

func (h *IngestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var doc milvus.Document

	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if doc.ID == 0 {
		http.Error(w, "id required", http.StatusBadRequest)
		return
	}

	if len(doc.Vector) == 0 {
		http.Error(w, "vector required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.mc.Insert(ctx, doc); err != nil {
		log.Println(err)
		http.Error(w, "insert failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})

}
