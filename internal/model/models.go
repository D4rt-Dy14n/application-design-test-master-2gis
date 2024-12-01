package model

import (
	"time"
)

type Order struct {
	ID        string    `json:"id"`
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
	CreatedAt time.Time `json:"created_at"`
}

type RoomAvailability struct {
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	Date    time.Time `json:"date"`
	Quota   int       `json:"quota"`
}

// New function for initial data
func InitialAvailabilityData() []RoomAvailability {
	return []RoomAvailability{
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
	}
}
