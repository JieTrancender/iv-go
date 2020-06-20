GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=bin
NAME=iv

# mac
build:
	CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME)/$(NAME)-mac

# linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)/$(NAME)-linux

# windows
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)/$(NAME)-win.exe

# 全平台
build-all:
	make build
	make build-win
	make build-linux

# 测试
test:
	go test -v ./ ./... -race -coverprofile=coverage.txt -covermode=atomic