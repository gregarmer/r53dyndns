DIST_DIR = dist

all: build

deps:
	go get "github.com/aws/aws-sdk-go/aws"
	go get "github.com/aws/aws-sdk-go/service/route53"
	go get "github.com/aws/aws-sdk-go/aws/session"
	go get "github.com/aws/aws-sdk-go/aws/credentials"

build: clean deps
	test -d $(DIST_DIR) || mkdir $(DIST_DIR)
	go build -o $(DIST_DIR)/r53dyndns main.go

clean:
	rm -rf $(DIST_DIR)

test: clean deps
	@go test -run=. -test.v ./config
	@go test -run=. -test.v ./dyndns
	@go test -run=. -test.v ./utils
