BUILD_DIR=./build
MODULE_NAME=pplate

.PHONY: setup
setup:
	dep ensure -v -vendor-only
.PHONY: test
test:
	go test -v ./...
.PHONY: build
build:
	go build -o ./build/$(MODULE_NAME) ./cmd/$(MODULE_NAME)/main.go
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)/*
.PHONY: gen
gen:
	go generate ./...
