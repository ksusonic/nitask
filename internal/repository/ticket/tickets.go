package ticket

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ksusonic/nitask/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	dbName         = "task"
	collectionName = "tickets"
)

func (s *Repository) List(ctx context.Context, in models.TicketListIn) ([]models.Ticket, error) {
	cur, err := s.Collection.Find(
		ctx,
		bson.M{"queue": in.Queue},
		options.Find().SetSort(bson.M{"createdAt": -1}),
		options.Find().SetLimit(in.Limit),
		options.Find().SetSkip(in.Offset),
	)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var tickets []models.Ticket
	if err = cur.All(ctx, &tickets); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (s *Repository) Get(ctx context.Context, key string) (*models.Ticket, error) {
	var ticket models.Ticket

	err := s.Collection.FindOne(ctx, bson.M{"key": key}).Decode(&ticket)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, models.ErrNotFound
	}

	return &ticket, err
}

func (s *Repository) Create(ctx context.Context, in models.TicketCreateIn) (*models.Ticket, error) {
	now := time.Now().UTC()

	var existing models.Ticket
	err := s.Collection.
		FindOne(ctx, bson.M{"idempotency_key": in.IdempotencyKey}).
		Decode(&existing)

	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		break
	case err == nil:
		return &existing, nil
	default:
		return nil, fmt.Errorf("check for existing ticket: %w", err)
	}

	key, err := s.createNextKey(ctx, in.Queue)
	if err != nil {
		return nil, fmt.Errorf("create next key: %w", err)
	}

	ticket := models.Ticket{
		Key:            key,
		Queue:          in.Queue,
		Title:          in.Title,
		Description:    in.Description,
		Status:         models.TicketStatusOpen,
		CreatedAt:      now,
		UpdatedAt:      now,
		IdempotencyKey: in.IdempotencyKey,
	}

	_, err = s.Collection.InsertOne(ctx, ticket)
	if err != nil {
		return nil, fmt.Errorf("insert ticket: %w", err)
	}

	return &ticket, nil
}

func (s *Repository) Update(ctx context.Context, key string, in models.TicketUpdateIn) (*models.Ticket, error) {
	update := bson.M{"updated_at": time.Now().UTC()}
	if in.Title != nil {
		update["title"] = *in.Title
	}
	if in.Description != nil {
		update["description"] = *in.Description
	}
	if in.Status != nil {
		update["status"] = *in.Status
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated models.Ticket
	err := s.Collection.FindOneAndUpdate(
		ctx,
		bson.M{"key": key},
		bson.M{"$set": update},
		opts,
	).Decode(&updated)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, models.ErrNotFound
	}
	return &updated, err
}

func (s *Repository) Delete(ctx context.Context, key string) error {
	res, err := s.Collection.DeleteOne(ctx, bson.M{"key": key})
	if res.DeletedCount == 0 {
		return models.ErrNotFound
	}

	return err
}
