package ticket

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type sequence struct {
	Seq int64 `bson:"seq"`
}

func (r *Repository) createNextKey(ctx context.Context, queueName string) (string, error) {
	var result sequence

	err := r.Collection.
		FindOneAndUpdate(
			ctx,
			bson.M{"_id": queueName},
			bson.M{"$inc": bson.M{"seq": 1}},
			options.FindOneAndUpdate().
				SetUpsert(true).
				SetReturnDocument(options.After),
		).
		Decode(&result)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%d", queueName, result.Seq), nil
}
