package milvus

import (
	"context"
	"log"
	"time"

	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/index"
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

func (c *Client) EnsureCollection(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Check whether the collection already exists.
	exists, err := c.cli.HasCollection(
		ctx,
		milvusclient.NewHasCollectionOption(CollectionName),
	)
	if err != nil {
		return err
	}

	// Create collection if it does not exist.
	if !exists {
		schema := entity.NewSchema().
			WithName(CollectionName).
			WithDescription("Semantic search documents").
			WithAutoID(false).
			WithField(
				entity.NewField().
					WithName(IDField).
					WithDataType(entity.FieldTypeInt64).
					WithIsPrimaryKey(true).
					WithIsAutoID(false),
			).
			WithField(
				entity.NewField().
					WithName(TextField).
					WithDataType(entity.FieldTypeVarChar).
					WithMaxLength(4096),
			).
			WithField(
				entity.NewField().
					WithName(VectorField).
					WithDataType(entity.FieldTypeFloatVector).
					WithDim(Dimension),
			)

		createOpt := milvusclient.NewCreateCollectionOption(
			CollectionName,
			schema,
		).WithIndexOptions(
			milvusclient.NewCreateIndexOption(
				CollectionName,
				VectorField,
				index.NewAutoIndex(entity.COSINE),
			),
		)

		if err := c.cli.CreateCollection(ctx, createOpt); err != nil {
			return err
		}
	}

	// Always load the collection.
	task, err := c.cli.LoadCollection(
		ctx,
		milvusclient.NewLoadCollectionOption(CollectionName),
	)
	if err != nil {
		return err
	}

	// Wait until the collection is fully loaded.
	if err := task.Await(ctx); err != nil {
		return err
	}

	return nil
}

func (c *Client) Insert(ctx context.Context, doc Document) error {
	return nil
}
