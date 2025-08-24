package storage

import "go.mongodb.org/mongo-driver/v2/mongo"

const (
	ticketDBName = "ticket"
)

func (s *Mongo) TicketDB() *mongo.Database {
	return s.client.Database(ticketDBName)
}
