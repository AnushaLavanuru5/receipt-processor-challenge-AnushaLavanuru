package main

import (
	"log"

	"receipt-processor/internal/handlers"
	"receipt-processor/internal/storage"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// In-memory storage
	store := storage.NewMemoryStore()

	// Registering the routes
	handler := handlers.NewReceiptHandler(store)
	app.Post("/receipts/process", handler.ProcessReceipt)
	app.Get("/receipts/:id/points", handler.GetReceiptPoints)

	// Starting the server
	log.Fatal(app.Listen(":8080"))
}
