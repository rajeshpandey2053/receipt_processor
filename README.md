# Receipt Processor API

The **Receipt Processor API** allows users to process receipts by submitting a receipt in JSON format and retrieving the number of points awarded based on predefined rules. This API is built with Go, uses a RESTful design, and supports containerization through Docker. Additionally, Swagger documentation is provided to explore and test the endpoints.

## Features

- **Process Receipts**: Submit a receipt and get a unique receipt ID.
- **Retrieve Points**: Retrieve the number of points awarded for a receipt using its ID.
- **Swagger Documentation**: API documentation and testing available through Swagger UI.
- **Dockerized Setup**: Fully containerized application for easy deployment.

## Table of Contents
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [API Endpoints](#api-endpoints )
- [Running Tests](#running-tests)

## Tech Stack

- **Language**: Go
- **Database**: In-memory (no persistent storage)
- **Containerization**: Docker
- **API Documentation**: Swagger
- **Testing**: Go testing framework, Testify


## Installation
### Clone the repository:

```bash
git clone https://github.com/rajeshpandey2053/fetch_receipt_processor.git
cd fetch_receipt_processor
```


### 2. Install Dependencies
You will need the following installed on your local machine
- docker -- [install guide](https://docs.docker.com/get-docker/)

### Build and Run using Docker
Make sure docker is running and then run docker compose to start the server
```bash
docker compose up --build
```

## API Endpoints
Use Swagger: The swagger endpoint is available at http://localhost:8080/swagger/index.html. You can test and execute the post and get request in the UI.

## Running Tests
To run the tests, you can use following command:

```bash
go test ./tests
```
