package models

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Points struct {
	Points int `json:"points"`
}

type ReceiptId struct {
	Id string `json:"id"`
}

// Response  structure for the result structure of ProcessReceipt.
type ProcessReceiptResponse struct {
	Id string `json:"id"`
}

// Response structure for the result structure of GetPointsById.
type GetPointsByIdResponse struct {
	Points int `json:"points"`
}
