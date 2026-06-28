# 🔎 Local RAG System (Go + ONNX + Milvus)

ระบบ Retrieval-Augmented Generation (RAG) แบบ local-first:
- Embedding รันด้วย ONNX (ไม่ต้องใช้ Python runtime)
- Vector search ด้วย Milvus
- API orchestrator ด้วย Go
- รองรับ LLM (OpenAI หรือ local llama.cpp)

---

# 🧠 Architecture

User Query
    ↓
Go API
    ↓
ONNX Embedding Runtime
    ↓
Milvus Vector DB
    ↓
Top-K Context
    ↓
LLM (OpenAI / local llama.cpp)
    ↓
Final Answer

---

# 🧱 Components

## Vector Database
Milvus
ใช้เก็บ:
- embedding vectors
- text chunks
- similarity index

## Embedding Engine
ONNX
- no Python runtime
- low latency inference
- deploy as binary

## LLM (Optional)
- OpenAI API
- llama.cpp (local)

## API Layer
- Go (Gin / Fiber / net/http)

---

# 🚀 Run Infrastructure

docker compose up -d

---

# ⚙️ Go Setup

go get github.com/milvus-io/milvus-sdk-go/v2
go get github.com/microsoft/onnxruntime-go
go get github.com/sugarme/tokenizer

---

# 🔎 Core Pipeline

Embedding → Milvus Search → Context → LLM

---

# 🧠 Summary

Local-first RAG system using Go + ONNX + Milvus


data set Semantic Textual Similarity hugging face