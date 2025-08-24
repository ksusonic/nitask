package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/ksusonic/nitask/pkg/api"
)

type TicketStatus api.TicketStatus

const (
	TicketStatusOpen       TicketStatus = TicketStatus(api.TicketStatusOpen)
	TicketStatusInProgress TicketStatus = TicketStatus(api.TicketStatusInProgress)
	TicketStatusDone       TicketStatus = TicketStatus(api.TicketStatusDone)
)

type Ticket struct {
	Key            string       `bson:"key"                   json:"key"`
	Queue          string       `bson:"queue"                 json:"queue"`
	Title          string       `bson:"title"                 json:"title"`
	Description    string       `bson:"description,omitempty" json:"description,omitempty"`
	Status         TicketStatus `bson:"status"                json:"status"`
	CreatedAt      time.Time    `bson:"created_at"            json:"created_at"`
	UpdatedAt      time.Time    `bson:"updated_at"            json:"updated_at"`
	IdempotencyKey uuid.UUID    `bson:"idempotency_key"       json:"idempotency_key"`
}

type TicketListIn struct {
	Queue  string
	Offset int64
	Limit  int64
	// TODO: add order/sort
}

type TicketCreateIn struct {
	Queue          string    `bson:"queue"                 json:"queue"`
	Title          string    `bson:"title"                 json:"title"`
	Description    string    `bson:"description,omitempty" json:"description,omitempty"`
	IdempotencyKey uuid.UUID `bson:"idempotency_key"       json:"idempotency_key"`
}

type TicketUpdateIn struct {
	Title       *string `bson:"title,omitempty"       json:"title,omitempty"`
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	Status      *string `bson:"status,omitempty"      json:"status,omitempty"`
}
