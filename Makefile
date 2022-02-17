.PHONY: build
build: 
		go build -v ./cmd/apiserver

.PHONY: test
test:
		go test -v -race -timeout 30s ./...

export HTTPS_PROXY=http://10.131.72.242:8080