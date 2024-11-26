package service

import (
	"fetch_receipt_processor/pkg/models"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

var pointsCollection = map[string]float64{
	"RuleAplhaNumeric": 1,
	"RuleDollarAmt": 50,
	"RuleTotalMultiple": 25,
	"RuleEveryItem": 5,
	"RuleItemDesc": 0.2, // This will be calculated dynamically based on the item price
	"RuleOddDate": 6,
	"RuleTime": 10,
}

// Rule defines the interface for a points calculation rule
type Rule interface {
	Calculate(receipt models.Receipt) int
}

// RuleAplhaNumeric implements Rule for alphanumeric characters in the retailer name
type RuleAplhaNumeric struct{}

func (r RuleAplhaNumeric) Calculate(receipt models.Receipt) int {
	points := 0
	for _, char := range receipt.Retailer {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			points += int(pointsCollection["RuleAplhaNumeric"])
		}
	}
	return points
}

// RuleDollarAmt implements Rule for round dollar amount total
type RuleDollarAmt struct{}

func (r RuleDollarAmt) Calculate(receipt models.Receipt) int {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		log.Printf("RuleDollarAmt: error parsing total: %v\n", err)
		return 0
	}
	if total == float64(int(total)) {
		return int(pointsCollection["RuleDollarAmt"])
	}
	return 0
}

// RuleTotalMultiple implements Rule for total being a multiple of 0.25
type RuleTotalMultiple struct{}

func (r RuleTotalMultiple) Calculate(receipt models.Receipt) int {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		log.Printf("RuleTotalMultiple: error parsing total: %v\n", err)
		return 0
	}
	if math.Mod(total, 0.25) == 0 {
		return int(pointsCollection["RuleTotalMultiple"])
	}
	return 0
}

// RuleEveryItem implements Rule for every two items on the receipt
type RuleEveryItem struct{}

func (r RuleEveryItem) Calculate(receipt models.Receipt) int {
	return (len(receipt.Items) / 2) * int(pointsCollection["RuleEveryItem"])
}

// RuleItemDesc implements Rule for item descriptions
type RuleItemDesc struct{}

func (r RuleItemDesc) Calculate(receipt models.Receipt) int {
	points := 0
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				log.Printf("RuleItemDesc: error parsing item price: %v\n", err)
				continue
			}
			points += int(math.Ceil(price * pointsCollection["RuleItemDesc"]))
		}
	}
	return points
}

// RuleOddDate implements Rule for odd day in the purchase date
type RuleOddDate struct{}

func (r RuleOddDate) Calculate(receipt models.Receipt) int {
	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		log.Printf("RuleOddDate: error parsing purchase date: %v\n", err)
		return 0
	}
	if date.Day()%2 != 0 {
		return int(pointsCollection["RuleOddDate"])
	}
	return 0
}

// RuleTime implements Rule for purchase time between 2:00pm and 4:00pm
type RuleTime struct{}

func (r RuleTime) Calculate(receipt models.Receipt) int {
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		log.Printf("RuleTime: error parsing purchase time: %v\n", err)
		return 0
	}
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		return int(pointsCollection["RuleTime"])
	}
	return 0
}
