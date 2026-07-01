package handler

import (
	"context"
	"net/http"
	"search-platform/internal/milvus"
	"time"
)

type ReadyHandler struct {
	mc *milvus.Client
}

func NewReadyHandler(mc *milvus.Client) *ReadyHandler {
	return &ReadyHandler{mc: mc}
}

func (h *ReadyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := h.mc.Ping(ctx); err != nil {
		writeJSONError(w, http.StatusServiceUnavailable, "milvus not ready")
		return
	}

	writeJSON(w, http.StatusOK, StatusResponse{Status: "ready"})
}
