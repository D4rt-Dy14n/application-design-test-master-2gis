package testdata

import (
	"applicationDesignTest/internal/model"
	"time"
)

func InitialAvailabilityData() []model.RoomAvailability {
	return []model.RoomAvailability{
		{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			Quota:   5,
		},
		{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			Quota:   5,
		},
		{
			HotelID: "reddison",
			RoomID:  "lux",
			Date:    time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC),
			Quota:   5,
		},
		{
			HotelID: "hilton",
			RoomID:  "deluxe",
			Date:    time.Date(2024, 1, 7, 0, 0, 0, 0, time.UTC),
			Quota:   5,
		},
		{
			HotelID: "hilton",
			RoomID:  "deluxe",
			Date:    time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC),
			Quota:   5,
		},
		{
			HotelID: "hilton",
			RoomID:  "deluxe",
			Date:    time.Date(2024, 1, 9, 0, 0, 0, 0, time.UTC),
			Quota:   5,
		},
	}
}
