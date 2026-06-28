package milvus

import (
	"context"
	"log"
	"time"

	"github.com/milvus-io/milvus/client/v2/milvusclient"
)

type Client struct {
	cli *milvusclient.Client
}

func New(address string) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cli, err := milvusclient.New(ctx, &milvusclient.ClientConfig{
		Address: address,
	})

	if err != nil {
		return nil, err
	}

	log.Println("milvus connected:", address)

	return &Client{cli: cli}, nil
}

func (c *Client) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	_, err := c.cli.ListCollections(ctx, milvusclient.NewListCollectionOption())
	return err
}
