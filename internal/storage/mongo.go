package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/ksusonic/nitask/pkg/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const connectTimeout = 3 * time.Second

type Mongo struct {
	client *mongo.Client
}

func NewMongo(config config.MongoDBConfig) (*Mongo, error) {
	client, err := mongo.Connect(
		options.Client().
			ApplyURI(config.URI).
			SetMaxPoolSize(config.MaxPoolSize).
			SetConnectTimeout(config.ConnectTimeout).
			SetCompressors([]string{"snappy", "zlib", "zstd"}),
	)
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("ping mongodb instance: %w", err)
	}

	return &Mongo{
		client: client,
	}, nil
}

func (s *Mongo) Client() *mongo.Client {
	return s.client
}

func (s *Mongo) Close(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}
