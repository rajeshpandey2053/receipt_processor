package handler

import (
	"encoding/json"
	"fetch_receipt_processor/pkg/models"
	"fetch_receipt_processor/pkg/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ReceiptHandler struct {
	service *service.ReceiptService
}

func NewReceiptHandler(service *service.ReceiptService) *ReceiptHandler {
	return &ReceiptHandler{service: service}
}

func StatusBadRequestErrorHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("The receipt is invalid"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("No receipt found for that id"))
}

// ProcessReceipt processes a receipt and returns the receipt ID.
// @Summary Process a receipt
// @Description Process the receipt and return the ID
// @Tags receipts
// @Accept json
// @Produce json
// @Param receipt body models.Receipt true "Receipt"
// @Success 200 {object} models.ProcessReceiptResponse
// @Failure 400 {string} string "The receipt is invalid"
// @Router /receipts/process [post]
func (h *ReceiptHandler) ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	log.Println("ReceiptHandler::ProcessReceipt: processing receipt")

	var receipt models.Receipt

	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		log.Printf("ReceiptHandler::ProcessReceipt: error decoding request body: %v\n", err)
		StatusBadRequestErrorHandler(w, r)
		return
	}

	log.Printf("ReceiptHandler::ProcessReceipt: received receipt: %+v\n", receipt)

	// Process receipt
	id, err := h.service.ProcessReceipt(receipt)
	if err != nil {
		log.Printf("ReceiptHandler::ProcessReceipt: error processing receipt: %v\n", err)
		StatusBadRequestErrorHandler(w, r)

		return
	}

	response := models.ProcessReceiptResponse{Id: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetPointsById retrieves the points for a receipt by its ID.
// @Summary Get points by receipt ID
// @Description Get the points associated with a specific receipt ID
// @Tags receipts
// @Param id path string true "Receipt ID"
// @Produce json
// @Success 200 {object} models.GetPointsByIdResponse
// @Failure 404 {string} string "No receipt found for that id"
// @Router /receipts/{id}/points [get]
func (h *ReceiptHandler) GetPointsById(w http.ResponseWriter, r *http.Request) {
	log.Println("ReceiptHandler::GetPointsById: getting points by id")

	vars := mux.Vars(r)
	id := vars["id"]

	points, err := h.service.GetPoints(id)
	if err != nil {
		log.Printf("ReceiptHandler::GetPointsById: error getting points: %v\n", err)
		NotFoundHandler(w, r)
		return
	}

	response := models.GetPointsByIdResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

