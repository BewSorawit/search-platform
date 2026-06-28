package main

import (
	"log"
	"net/http"
	"search-platform/internal/milvus"
	"search-platform/internal/router"
)

func main() {

	// connect milvus
	mc, err := milvus.New("localhost:19530")
	if err != nil {
		log.Fatal("milvus connection failed:", err)
	}

	// inject dependency
	r := router.New(mc)

	log.Println("Server started on :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
