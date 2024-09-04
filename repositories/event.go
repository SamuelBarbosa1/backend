package repositories

import (
	"context"
	"time"

	"github.com/SamuelBarbosa1/ticket-booking-project-v1/models"
)

type EventRepository struct {
	db any
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	events = append(events, &models.Event{
		ID:        "023497832984785238974",
		Name:      "Minha Banda Favorita",
		Location:  "Em algum lugar dali",
		Date:      time.Now(),
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	})

	return events, nil
}
func (r *EventRepository) GetOne(ctx context.Context, eventId uint) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {
	return nil, nil
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
