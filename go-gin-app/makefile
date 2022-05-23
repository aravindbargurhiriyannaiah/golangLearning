BINARY_NAME=go-gin-app
artifacts=./artifacts

build:
	go build -o ${artifacts}/${BINARY_NAME} httpd/main.go

run:
	./${artifacts}/${BINARY_NAME}

build_and_run: build run

test: 
	go test ./...

test_coverage:
	go test ./... -coverprofile=./${artifacts}/${BINARY_NAME}_test_coverage.out

format:
	go fmt ./...
