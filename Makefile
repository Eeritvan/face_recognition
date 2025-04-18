GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=face_recognition
SRC_DIR=./src

.PHONY: start build test clean

start: build
	cd $(SRC_DIR) && ./$(BINARY_NAME) $(ARGS)

build:
	$(GOBUILD) -C $(SRC_DIR) -o $(BINARY_NAME)

test: build
	$(GOTEST) -C $(SRC_DIR) $(shell cat $(SRC_DIR)/testdirs.txt) -cover

clean:
	cd $(SRC_DIR) && rm -f $(BINARY_NAME)