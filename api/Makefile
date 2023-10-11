OUTDIR=build
BIN=storageapi

GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean

build:
	$(GOBUILD) -o $(OUTDIR)/$(BIN) ./cmd/storage-api

run:
	./$(OUTDIR)/$(BIN)

clean:
	$(GOCLEAN)
	-rm -r build

.PHONY: build run clean
