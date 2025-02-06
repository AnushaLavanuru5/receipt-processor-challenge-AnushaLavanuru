package handlers

import (
	"receipt-processor/internal/models"
	"receipt-processor/internal/services"
	"receipt-processor/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ReceiptHandler struct {
	store *storage.MemoryStore
}

func NewReceiptHandler(store *storage.MemoryStore) *ReceiptHandler {
	return &ReceiptHandler{store: store}
}

func (h *ReceiptHandler) ProcessReceipt(c *fiber.Ctx) error {
	var receipt models.Receipt
	if err := c.BodyParser(&receipt); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	id := uuid.New().String()
	h.store.Save(id, receipt)

	return c.JSON(fiber.Map{"id": id})
}

func (h *ReceiptHandler) GetReceiptPoints(c *fiber.Ctx) error {
	id := c.Params("id")
	receipt, exists := h.store.Get(id)
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Receipt not found"})
	}

	points, err := services.CalculatePoints(receipt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"points": points})
}
