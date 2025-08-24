package ticket

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	collectionName = "tickets"
)

type Repository struct {
	Collection *mongo.Collection
}

func NewRepository(ctx context.Context, database *mongo.Database, withEnsureIndexes bool) (*Repository, error) {
	collection := database.Collection(collectionName)

	if withEnsureIndexes {
		if err := ensureIndexes(ctx, collection); err != nil {
			return nil, fmt.Errorf("failed to ensure indexes: %w", err)
		}
	}

	return &Repository{
		Collection: collection,
	}, nil
}
