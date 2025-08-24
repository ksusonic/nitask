package storage

import (
	"github.com/ksusonic/nitask/internal/storage/tickets"
)

const (
	tasksDB               = "tasks"
	ticketsCollectionName = "tickets"
)

type databases struct {
	tickets *tickets.Collection
}

func (s *Mongo) TicketsCollection() *tickets.Collection {
	if s.databases.tickets == nil {
		collection := s.client.Database(tasksDB).Collection(ticketsCollectionName)

		s.databases.tickets = tickets.NewCollection(collection)
	}

	return s.databases.tickets
}
