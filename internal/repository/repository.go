package repository

import (
	"context"
	"sync"
	"time"

	"applicationDesignTest/internal/model"
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
	mutex        sync.RWMutex
	availability []model.RoomAvailability
	orders       []model.Order
}

func NewMemoryRepository(initialAvailability []model.RoomAvailability) *MemoryRepository {
	return &MemoryRepository{
		availability: initialAvailability,
		orders:       make([]model.Order, 0),
	}
}

func (r *MemoryRepository) GetAvailability(ctx context.Context, hotelID, roomID string, date time.Time) (*model.RoomAvailability, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		r.mutex.RLock()
		defer r.mutex.RUnlock()

		for _, a := range r.availability {
			if a.HotelID == hotelID && a.RoomID == roomID && a.Date.Equal(date) {
				return &a, nil
			}
		}
		return nil, nil
	}
}
func (r *MemoryRepository) UpdateAvailability(ctx context.Context, availability *model.RoomAvailability) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		r.mutex.Lock()
		defer r.mutex.Unlock()

		for i, a := range r.availability {
			if a.HotelID == availability.HotelID && a.RoomID == availability.RoomID && a.Date.Equal(availability.Date) {
				r.availability[i] = *availability
				return nil
			}
		}
		r.availability = append(r.availability, *availability)
		return nil
	}
}

func (r *MemoryRepository) CreateOrder(ctx context.Context, order *model.Order) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		r.mutex.Lock()
		defer r.mutex.Unlock()

		r.orders = append(r.orders, *order)
		return nil
	}
}

func (r *MemoryRepository) GetOrders(ctx context.Context) ([]model.Order, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		r.mutex.RLock()
		defer r.mutex.RUnlock()

		orders := make([]model.Order, len(r.orders))
		copy(orders, r.orders)
		return orders, nil
	}
}
