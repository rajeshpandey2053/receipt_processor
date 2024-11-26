package main

import (
	"fetch_receipt_processor/pkg/database"
	"fetch_receipt_processor/pkg/handler"
	"fetch_receipt_processor/pkg/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "fetch_receipt_processor/docs"

	"github.com/spf13/viper"
)

// @title Receipt Processor API
// @version 1.0
// @description This is an api developed for Fetch Rewards receipt processor challenge.
// @host localhost:8080
// @BasePath /
func main(){
	setupConfig()

	receiptDb := database.NewInMemoryDb()
	receitService := service.NewReceiptService(receiptDb)
	receiptHandler := handler.NewReceiptHandler(receitService)

	router := mux.NewRouter()

	receiptsRouter := router.PathPrefix("/receipts").Subrouter()
	receiptsRouter.HandleFunc("/process", receiptHandler.ProcessReceipt).Methods("POST")
	receiptsRouter.HandleFunc("/{id}/points", receiptHandler.GetPointsById).Methods("GET")

	// Swagger documentation route
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	addr := ":" + viper.GetString("server.port")
	log.Println("Server started on", addr)
	log.Fatal(http.ListenAndServe(addr, router))

}

// setupConfig loads the configuration from a file
func setupConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}