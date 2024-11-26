package service

import (
	"fetch_receipt_processor/pkg/database"
	"fetch_receipt_processor/pkg/models"
	"log"
)

// ReceiptService is the service layer for receipt processing.
type ReceiptService struct {
	db database.Database
}

// NewReceiptService creates a new ReceiptService instance.
func NewReceiptService(db database.Database) *ReceiptService {
	return &ReceiptService{db: db}
}

// CalculatePoints calculates the points for a given receipt.
func (s *ReceiptService) CalculatePoints(receipt models.Receipt) int {
	log.Println("ReceiptService::CalculatePoints: calculating points for receipt")
	points := 0
	rules := []Rule{
		RuleAplhaNumeric{},
		RuleDollarAmt{},
		RuleTotalMultiple{},
		RuleEveryItem{},
		RuleItemDesc{},
		RuleOddDate{},
		RuleTime{},
	}

	for _, rule := range rules {
		points += rule.Calculate(receipt)
	}
	return points
}


// GetPointsById retrieves points associated with a receipt ID.
func (s *ReceiptService) ProcessReceipt(receipt models.Receipt) (string, error) {
	log.Println("ReceiptService::GetPointsById: Processing started")
	points := s.CalculatePoints(receipt)
	id, err := s.db.AddPoints(points)
	if err != nil {
		log.Println("ReceiptService::ProcessReceipt: failed to insert receipt")
		return "", err
	}
	log.Printf("ReceiptService::ProcessReceipt: successfully inserted receipt with ID %s\n", id)
	return id, err
}

// GetPointsById retrieves points associated with a receipt ID.
func (s *ReceiptService) GetPoints(id string) (int, error) {
	log.Printf("ReceiptService::GetPointsById: retrieving points for receipt")
	points, err := s.db.GetPointsById(id)
	if err != nil {
		log.Printf("ReceiptService::GetPointsById: failed to retrieve points for recetipt ID %s: %v\n", id, err)
		return -1, err
	}
	log.Printf("ReceiptService::GetPointsById: successfully retrieved points for receipt ID %s\n", id)
	return points, nil
}

