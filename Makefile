GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

.PHONY: build
build:
	$(GOBUILD) -o mongocli -v mongocli.go

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: test/e2e
test/e2e:
	docker build -f Dockerfile -t mongocli:latest .
	docker-compose -f docker-compose.yml up --abort-on-container-exit --exit-code-from app
