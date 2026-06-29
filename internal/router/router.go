package router

import (
	"net/http"
	"search-platform/internal/handler"
	"search-platform/internal/milvus"
)

func New(mc *milvus.Client) *http.ServeMux {
	mux := http.NewServeMux()

	// liveness
	mux.HandleFunc("/health", handler.Health)

	mux.Handle("/ready", handler.NewReadyHandler(mc))

	mux.Handle("/ingest", handler.NewIngestHandler(mc))

	return mux
}
