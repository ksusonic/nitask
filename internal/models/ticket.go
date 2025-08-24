package models

import "time"

type Ticket struct {
	ID        string    `bson:"_id"`
	Title     string    `bson:"title"`
	Content   string    `bson:"content"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
