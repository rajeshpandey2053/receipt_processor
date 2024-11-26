# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=main
MAIN=cmd/main.go


# Swagger parameters
SWAGGER_OUT=docs
SWAGGER_DIR=cmd,pkg

# Targets
all: swagger build

build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN)

test:
	$(GOTEST) -v ./tests/...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GORUN)  $(MAIN)

swagger:
	swag init  -d $(SWAGGER_DIR) -o $(SWAGGER_OUT)

.PHONY: all build test clean run swagger