PHONY: build-binary
build-binary: clean fmt
	rm -f go.sum
	go get -v -t -d ./...
	$(eval DIR := $(GOPATH)/src/examples/ch11)
	CGO_ENABLED=0 GO111MODULE=on GOOS=$(GOOS) go build -o $(GOPATH)/bin/mychannel
	go mod tidy
	rm -f go.sum
	
PHONY: fmt
fmt:
	gofmt -w *.go
	

PHONY: clean
clean:
	rm -f $(GOPATH)/bin/mychannel
	rm -f go.sum
	
