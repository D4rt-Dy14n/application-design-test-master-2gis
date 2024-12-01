package repository

import (
    "context"
    "sync"
    "time"

    "github.com/d4rt-dy14n/application-design-test-master/internal/model"
)

type AvailabilityRepository interface {
    GetAvailability(ctx context.Context, hotelID, roomID string, date time.Time) (*model.RoomAvailability, error)
    UpdateAvailability(ctx context.Context, availability *model.RoomAvailability) error
}

type OrderRepository interface {
    CreateOrder(ctx context.Context, order *model.Order) error
    GetOrders(ctx context.Context) ([]model.Order, error)
}

type MemoryRepository struct {
    mutex         sync.RWMutex
    availability  []model.RoomAvailability
    orders        []model.Order
}

func NewMemoryRepository(initialAvailability []model.RoomAvailability) *MemoryRepository {
    return &MemoryRepository{
        availability: initialAvailability,
        orders:      make([]model.Order, 0),
    }
}
