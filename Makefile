SHELL = /bin/bash
DIST_DIR = dist

.PHONY: build

all: build

deps:
	go get "github.com/aws/aws-sdk-go/aws"
	go get "github.com/aws/aws-sdk-go/service/route53"
	go get "github.com/aws/aws-sdk-go/aws/session"
	go get "github.com/aws/aws-sdk-go/aws/credentials"
	go get "github.com/miekg/dns"
	test -d $(DIST_DIR) || mkdir $(DIST_DIR)

freebsd:
	@echo [go] building FreeBSD amd64 binary...
	@time GOOS=freebsd GOARCH=amd64 go build -o $(DIST_DIR)/r53dyndns.fbsd main.go

openbsd:
	@echo [go] building OpenBSD amd64 binary...
	@time GOOS=openbsd GOARCH=amd64 go build -o $(DIST_DIR)/r53dyndns.obsd main.go

linux:
	@echo [go] building Linux amd64 binary...
	@time GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/r53dyndns.linux main.go

build: clean deps freebsd openbsd linux

clean:
	rm -rf $(DIST_DIR)

test: clean deps
	@go test -run=. -test.v ./config
	@go test -run=. -test.v ./dyndns
	@go test -run=. -test.v ./utils
