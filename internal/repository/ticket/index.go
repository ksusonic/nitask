package ticket

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ensureIndexes(ctx context.Context, collection *mongo.Collection) error {
	_, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    bson.M{"key": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{"idempotency_key": 1},
			Options: options.Index().SetUnique(true).
				SetPartialFilterExpression(bson.M{"idempotency_key": bson.M{"$exists": true}}),
		},
		{
			Keys:    bson.M{"queue": 1},
			Options: options.Index(),
		},
	})

	return err
}
