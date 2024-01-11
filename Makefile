TAG 				?= kind
IMG_NAME   	:= gophercise-choose-your-own-adventure
IMAGE      	:= ${IMG_NAME}:${TAG}

default: help

test:   ## Run all tests
	@go clean --testcache && go test ./... -v

cover:  ## Run test coverage suite
	@go test ./... --coverprofile=cov.out
	@go tool cover --html=cov.out

img:    ## Build Docker Image
	@docker build --rm -t ${IMAGE} .

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'

lint:
	@docker run -t --rm -v .:/app -w /app golangci/golangci-lint:v1.55.2 golangci-lint run -v

dev: 			## Run the application in development mode
	@go run .

vet:			## Run go vet
	@go vet -json
