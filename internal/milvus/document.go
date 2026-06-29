package milvus

type Document struct {
	ID     int64     `json:"id"`
	Text   string    `json:"text"`
	Vector []float32 `json:"vector"`
}

