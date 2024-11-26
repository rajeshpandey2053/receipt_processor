package tests

import (
	"bytes"
	"encoding/json"
	"fetch_receipt_processor/pkg/database"
	"fetch_receipt_processor/pkg/handler"
	"fetch_receipt_processor/pkg/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper function to set up a router
func setupRouter() (*mux.Router, *database.InMemoryDb) {

	mockreceiptDb := database.NewInMemoryDb()
	mockreceitService := service.NewReceiptService(mockreceiptDb)
	mockreceiptHandler := handler.NewReceiptHandler(mockreceitService)

	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", mockreceiptHandler.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", mockreceiptHandler.GetPointsById).Methods("GET")
	return router, &mockreceiptDb
}


func TestReceiptHandler_ProcessReceipt(t *testing.T) {
	router,_ := setupRouter()

	tests := []struct {
		name           string
		payload        []byte
		expectedStatus int
		expectedBody   string
		checkJSONField string
	}{
		{
			name: "Valid Receipt",
			payload: []byte(`{
				"retailer":"Walmart",
				"purchaseDate":"2024-06-26",
				"purchaseTime":"10:25",
				"items":[{"shortDescription":"xyz","price":"10.01"}],
				"total":"10.01"
			}`),
			expectedStatus: http.StatusOK,
			checkJSONField: "id",
		},
		{
			name:           "Invalid JSON",
			payload:        []byte(`{"invalid_json":`),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "The receipt is invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(tt.payload))
			req.Header.Set("Content-Type", "application/json")
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedBody != "" {
				assert.Equal(t, tt.expectedBody, rr.Body.String())
			}

			if tt.checkJSONField != "" {
				var responseBody map[string]interface{}
				err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
				
				require.NoError(t, err, "Response body should be valid JSON")
				_, exists := responseBody[tt.checkJSONField]
				assert.True(t, exists, "Response should contain a '%s' field", tt.checkJSONField)
			}
		})
	}
}

func TestReceiptHandler_GetPointsByID(t *testing.T) {
	router, mockreceiptDb := setupRouter()

	testID, _ := mockreceiptDb.AddPoints(100)
	println(testID)

	tests := []struct {
		name           string
		id             string
		expectedStatus int
		expectedBody   string
		checkJSONField string
	}{
		{
			name:           "Valid ID",
			id:             testID,
			expectedStatus: http.StatusOK,
			checkJSONField: "points",
		},
		{
			name:           "Invalid ID",
			id:             "invalidID",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "No receipt found for that id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/receipts/"+tt.id+"/points", nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedBody != "" {
				assert.Equal(t, tt.expectedBody, rr.Body.String())
			}

			if tt.checkJSONField != "" {
				var responseBody map[string]interface{}
				err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
				require.NoError(t, err, "Response body should be valid JSON")
				_, exists := responseBody[tt.checkJSONField]
				assert.True(t, exists, "Response should contain a '%s' field", tt.checkJSONField)
			}
		})
	}
}
