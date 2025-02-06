package test

import (
	"receipt-processor/internal/models"
	"receipt-processor/internal/services"
	"testing"
)

func TestCalculatePoints_Example1(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
	}

	points, err := services.CalculatePoints(receipt)
	if err != nil {
		t.Fatalf("Error calculating points: %v", err)
	}

	expectedPoints := 28
	if points != expectedPoints {
		t.Errorf("Expected %d points, got %d", expectedPoints, points)
	}
}

func TestCalculatePoints_Example2(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Total:        "9.00",
		Items: []models.Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
	}

	points, err := services.CalculatePoints(receipt)
	if err != nil {
		t.Fatalf("Error calculating points: %v", err)
	}

	expectedPoints := 109
	if points != expectedPoints {
		t.Errorf("Expected %d points, got %d", expectedPoints, points)
	}
}
