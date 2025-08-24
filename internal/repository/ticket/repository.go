package ticket

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	Collection *mongo.Collection
}

func NewRepository(ctx context.Context, client *mongo.Client, withEnsureIndexes bool) (*Repository, error) {
	collection := client.Database(dbName).Collection(collectionName)

	if withEnsureIndexes {
		if err := ensureIndexes(ctx, collection); err != nil {
			return nil, fmt.Errorf("failed to ensure indexes: %w", err)
		}
	}

	return &Repository{
		Collection: collection,
	}, nil
}
