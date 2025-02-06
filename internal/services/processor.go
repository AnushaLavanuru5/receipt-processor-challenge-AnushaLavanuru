package services

import (
	"math"
	"strconv"
	"strings"
	"time"

	"receipt-processor/internal/models"
)

func CalculatePoints(receipt models.Receipt) (int, error) {
	points := 0

	for _, char := range receipt.Retailer {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			points++
		}
	}

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, err
	}

	if total == float64(int(total)) {
		points += 50
	}

	if int(total*100)%25 == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	parsedDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return 0, err
	}
	if parsedDate.Day()%2 != 0 {
		points += 6
	}

	parsedTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}
	if parsedTime.Hour() >= 14 && parsedTime.Hour() < 16 {
		points += 10
	}

	return points, nil
}
