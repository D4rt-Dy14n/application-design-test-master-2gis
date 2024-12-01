package main

import (
	"applicationDesignTest/internal/handler"
	"applicationDesignTest/internal/repository"
	"applicationDesignTest/internal/service"
	"applicationDesignTest/internal/testdata"
	"log"
	"net/http"
)

func main() {
	memoryRepo := repository.NewMemoryRepository(testdata.InitialAvailabilityData())
	bookingService := service.NewBookingService(memoryRepo, memoryRepo)
	bookingHandler := handler.NewBookingHandler(bookingService)

	mux := http.NewServeMux()
	mux.HandleFunc("/orders", bookingHandler.CreateOrder)

	log.Println("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed: ", err)
	}
}
