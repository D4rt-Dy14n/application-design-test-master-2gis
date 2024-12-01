package service

import (
	"applicationDesignTest/internal/model"
	"applicationDesignTest/internal/repository"
	"context"
	"errors"
	"time"
)

var (
	ErrInvalidDateRange = errors.New("invalid date range")
	ErrNoAvailability   = errors.New("no availability for selected dates")
)

type BookingService struct {
	availabilityRepo repository.AvailabilityRepository
	orderRepo        repository.OrderRepository
}

func NewBookingService(ar repository.AvailabilityRepository, or repository.OrderRepository) *BookingService {
	return &BookingService{
		availabilityRepo: ar,
		orderRepo:        or,
	}
}

func getDaysBetween(start, end time.Time) []time.Time {
	var days []time.Time
	for d := start; d.Before(end) || d.Equal(end); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}
	return days
}

func (s *BookingService) CreateOrder(ctx context.Context, order *model.Order) error {
	if order.From.After(order.To) {
		return ErrInvalidDateRange
	}

	if order.ID == "" {
		order.ID = model.GenerateOrderID()
	}

	days := getDaysBetween(order.From, order.To)

	// Check availability for all days
	for _, day := range days {
		availability, err := s.availabilityRepo.GetAvailability(ctx, order.HotelID, order.RoomID, day)
		if err != nil || availability.Quota < 1 {
			return ErrNoAvailability
		}
	}

	// Update availability
	for _, day := range days {
		availability, _ := s.availabilityRepo.GetAvailability(ctx, order.HotelID, order.RoomID, day)
		availability.Quota--
		if err := s.availabilityRepo.UpdateAvailability(ctx, availability); err != nil {
			return err
		}
	}

	return s.orderRepo.CreateOrder(ctx, order)
}
