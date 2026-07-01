package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"search-platform/internal/milvus"
	"time"
)

type documentUpserter interface {
	Upsert(ctx context.Context, doc milvus.Document) error
}

type IngestHandler struct {
	mc documentUpserter
}

func NewIngestHandler(mc documentUpserter) *IngestHandler {
	return &IngestHandler{mc: mc}
}

func (h *IngestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var doc milvus.Document

	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if doc.ID < 0 {
		writeJSONError(w, http.StatusBadRequest, "id must be non-negative")
		return
	}

	if doc.Text == "" {
		writeJSONError(w, http.StatusBadRequest, "text required")
		return
	}

	if len(doc.Text) > 4096 {
		writeJSONError(w, http.StatusBadRequest, "text too long")
		return
	}

	if len(doc.Vector) != milvus.Dimension {
		writeJSONError(w, http.StatusBadRequest, "invalid vector dimension")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.mc.Upsert(ctx, doc); err != nil {
		status, msg := milvus.HTTPStatusAndMessage(err)
		if status >= http.StatusInternalServerError {
			log.Println(err)
		}
		writeJSONError(w, status, msg)
		return
	}

	writeJSON(w, http.StatusOK, StatusResponse{Status: "ok"})
}
