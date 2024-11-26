package database

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

// Database Interface
type Database interface {
	AddPoints(points int) (string, error)
	GetPointsById(id string) (int, error)
}

// InMemoryDb is the data repository for storing receipt points.
type InMemoryDb struct {
	receiptPoints map[string]int
}

// NewInMemoryDb creates a new InMemoryDb instance.
func NewInMemoryDb() InMemoryDb {
	return InMemoryDb{receiptPoints: make(map[string]int),}
	}

// GetPointsById retrieves points associated with a receipt ID.
func (db InMemoryDb) GetPointsById(id string) (int, error) {
	points, exists := db.receiptPoints[id]
	if !exists {
		log.Printf("Database::GetPointsById: receipt ID %s not found\n", id)
		return -1, errors.New("receipt Id not found")
	}
	log.Printf("Database::GetPointsById: successfully retrieved points for receipt ID %s\n", id)
	return points, nil
}

// AddPoints adds a new receipt entry with points and returns the receipt ID.
func (db InMemoryDb) AddPoints(points int) (string, error) {
	log.Println("Database::AddPoints: adding a new receipt entry")
	id := uuid.New().String()
	db.receiptPoints[id] = points
	log.Printf("Database::AddPoints: successfully added receipt entry with ID %s\n", id)
	return id, nil
}