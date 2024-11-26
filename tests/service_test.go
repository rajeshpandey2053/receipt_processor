package tests

import (
	"fetch_receipt_processor/pkg/database"
	"fetch_receipt_processor/pkg/models"
	"fetch_receipt_processor/pkg/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePoints(t *testing.T) {
	mockdb := database.NewInMemoryDb()
	mockreceiptService := service.NewReceiptService(mockdb)

	tests := []struct {
		name     string
		receipt  models.Receipt
		expected int
	}{
		{
			name: "Single item",
			receipt: models.Receipt{
				Retailer:     "Walgreens",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "08:13",
				Items: []models.Item{
					{
						ShortDescription: "Pepsi - 12-oz",
						Price:            "1.25",
					},
					{
						ShortDescription: "Dasani",
						Price:            "1.40",
					},
				},
				Total: "2.65",
			},
			expected: 15,
		},
		{
			name: "Single item from Target",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "13:13",
				Items: []models.Item{
					{
						ShortDescription: "Pepsi - 12-oz",
						Price:            "1.25",
					},
				},
				Total: "1.25",
			},
			expected: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := mockreceiptService.CalculatePoints(tt.receipt)
			assert.Equal(t, tt.expected, points)
		})
	}
}