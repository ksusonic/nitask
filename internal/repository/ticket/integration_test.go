//go:build integration

package ticket_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ksusonic/nitask/internal/models"
	"github.com/ksusonic/nitask/internal/repository/ticket"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	testMongoURI = "mongodb://localhost:27017"
	testDBName   = "integration_test_task_db"
	testQueue    = "testqueue"
)

func setupTestRepo(t *testing.T) (*ticket.Repository, func()) {
	t.Helper()

	client, err := mongo.Connect(options.Client().ApplyURI(testMongoURI))
	if err != nil {
		t.Fatalf("failed to connect to mongo: %v", err)
	}

	repo, err := ticket.NewRepository(
		t.Context(),
		client.Database(testDBName),
		true,
	)
	require.NoError(t, err)

	cleanup := func() {
		_ = client.Database(testDBName).Drop(t.Context())
		_ = client.Disconnect(t.Context())
	}
	return repo, cleanup
}

func TestRepository(t *testing.T) {
	repo, cleanup := setupTestRepo(t)
	defer cleanup()

	testUUID := uuid.New()

	t.Run("create_get", func(t *testing.T) {
		in := models.TicketCreateIn{
			Queue:          testQueue,
			Title:          "Test Ticket",
			Description:    "Test Description",
			IdempotencyKey: testUUID,
		}
		ticket, err := repo.Create(t.Context(), in)
		require.NoError(t, err)
		require.NotNil(t, ticket)
		require.NotEmpty(t, ticket.Key)

		got, err := repo.Get(t.Context(), ticket.Key)
		require.NoError(t, err)
		require.NotNil(t, got)

		ticket.CreatedAt = got.CreatedAt
		ticket.UpdatedAt = got.UpdatedAt

		require.Equal(t, ticket, got)
	})

	t.Run("get_all", func(t *testing.T) {
		// same idempotency key
		in := models.TicketCreateIn{
			Queue:          testQueue,
			Title:          "Test Ticket",
			Description:    "Test Description",
			IdempotencyKey: testUUID,
		}
		ticket, err := repo.Create(t.Context(), in)
		require.NoError(t, err)
		require.NotNil(t, ticket)
		require.NotEmpty(t, ticket.Key)

		got, err := repo.List(t.Context(), models.TicketListIn{
			Queue: testQueue,
		})
		require.NoError(t, err)
		require.NotNil(t, got)

		require.Equal(t, []models.Ticket{
			*ticket,
		}, got)

		ticket2, err := repo.Create(t.Context(), models.TicketCreateIn{
			Queue:          testQueue,
			Title:          "Test Ticket",
			Description:    "Test Description",
			IdempotencyKey: uuid.New(),
		})
		require.NoError(t, err)
		require.NotNil(t, ticket2)
		require.NotEmpty(t, ticket2.Key)
		require.NotEqual(t, ticket.Key, ticket2.Key)
	})

	t.Run("update_and_delete", func(t *testing.T) {
		in := models.TicketCreateIn{
			Queue:          testQueue,
			Title:          "ToUpdate",
			Description:    "Desc",
			IdempotencyKey: uuid.New(),
		}
		ticket, err := repo.Create(t.Context(), in)
		require.NoError(t, err)
		require.NotNil(t, ticket)
		require.NotEmpty(t, ticket.Key)

		newTitle := "Updated"
		updateIn := models.TicketUpdateIn{Title: &newTitle}
		updated, err := repo.Update(t.Context(), ticket.Key, updateIn)
		require.NoError(t, err)

		if updated.Title != newTitle {
			t.Errorf("expected updated title %q, got %q", newTitle, updated.Title)
		}

		err = repo.Delete(t.Context(), ticket.Key)
		if err != nil {
			t.Fatalf("Delete failed: %v", err)
		}
		_, err = repo.Get(t.Context(), ticket.Key)
		if err == nil {
			t.Error("expected error after delete, got nil")
		}
	})
}
