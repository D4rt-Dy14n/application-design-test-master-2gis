package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"applicationDesignTest/internal/model"
	"applicationDesignTest/internal/service"
)

type BookingHandler struct {
	bookingService *service.BookingService
}

func NewBookingHandler(bs *service.BookingService) *BookingHandler {
	return &BookingHandler{
		bookingService: bs,
	}
}

func (h *BookingHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.bookingService.CreateOrder(r.Context(), &order); err != nil {
		switch err {
		case service.ErrInvalidDateRange:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case service.ErrNoAvailability:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
	log.Printf("Order created: \n"+
		"ID: %s\n"+
		"Hotel: %s\n"+
		"Room: %s\n"+
		"Guest Email: %s\n"+
		"Check-in: %s\n"+
		"Check-out: %s\n",
		order.ID,
		order.HotelID,
		order.RoomID,
		order.UserEmail,
		order.From.Format("2006-01-02"),
		order.To.Format("2006-01-02"))
}
