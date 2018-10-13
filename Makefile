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
	@mkdir -p $(DIST_DIR)/freebsd
	@time GOOS=freebsd GOARCH=386 go build -o $(DIST_DIR)/freebsd/r53dyndns.386 main.go
	@time GOOS=freebsd GOARCH=amd64 go build -o $(DIST_DIR)/freebsd/r53dyndns.amd64 main.go

openbsd:
	@echo [go] building OpenBSD amd64 binary...
	@mkdir -p $(DIST_DIR)/openbsd
	@time GOOS=openbsd GOARCH=386 go build -o $(DIST_DIR)/openbsd/r53dyndns.386 main.go
	@time GOOS=openbsd GOARCH=amd64 go build -o $(DIST_DIR)/openbsd/r53dyndns.amd64 main.go
	@time GODEBUG=netdns=cgo GOOS=openbsd GOARCH=amd64 go build -o $(DIST_DIR)/openbsd/r53dyndns.amd64-cgo main.go

linux:
	@echo [go] building Linux amd64 binary...
	@mkdir -p $(DIST_DIR)/linux
	@time GOOS=linux GOARCH=386 go build -o $(DIST_DIR)/linux/r53dyndns.386 main.go
	@time GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/linux/r53dyndns.amd64 main.go

build: clean deps freebsd openbsd linux

clean:
	rm -rf $(DIST_DIR)

test: clean deps
	@go test -run=. -test.v ./config
	@go test -run=. -test.v ./dyndns
	@go test -run=. -test.v ./utils
