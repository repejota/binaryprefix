build: lint
	go build ./...

test:
	go test -v ./...

cover:
	go test -cover

lint:
	go vet ./...
	golint ./...

dev-deps:
	go get -u github.com/golang/lint/golint
