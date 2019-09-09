GO_VERSION = 1.13

.PHONY: all install
.DEFAULT_GOAL = all

all:
	@go get -d -v ./...
	@CGO_ENABLED=0 go build -ldflags '-w -s' -v

install:
	@go get -d -v ./...
	@CGO_ENABLED=0 go install -ldflags '-w -s' -v
