package tests

import (
	"fetch_receipt_processor/pkg/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryDb_AddPoints(t *testing.T) {
	mockdb := database.NewInMemoryDb()
	id, err := mockdb.AddPoints(100)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func TestInMemoryDb_GetPointsById(t *testing.T) {
	mockdb := database.NewInMemoryDb()
	id, err := mockdb.AddPoints(100)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	points, err := mockdb.GetPointsById(id)
	assert.NoError(t, err)
	assert.Equal(t, 100, points)
}

func TestInMemoryDb_GetPointsById_NotFound(t *testing.T) {
	mockdb := database.NewInMemoryDb()
	_, err := mockdb.GetPointsById("invalid_id")
	assert.Error(t, err)
}

