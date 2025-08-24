package tickets

import (
	"context"

	"github.com/ksusonic/nitask/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Collection struct {
	*mongo.Collection
}

func NewCollection(collection *mongo.Collection) *Collection {
	return &Collection{
		Collection: collection,
	}
}

func (c *Collection) FindOne(id string) (*models.Ticket, error) {
	var ticket models.Ticket
	err := c.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&ticket)
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}
