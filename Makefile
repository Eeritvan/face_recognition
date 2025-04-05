GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=face_recognition
SRC_DIR=./src

.PHONY: start build

start: build
	cd $(SRC_DIR) && ./$(BINARY_NAME) $(ARGS)

build:
	$(GOBUILD) -C $(SRC_DIR) -o $(BINARY_NAME)

test:
	$(GOTEST) -C $(SRC_DIR) ./... -cover

clean:
	cd $(SRC_DIR) && rm -f $(BINARY_NAME)