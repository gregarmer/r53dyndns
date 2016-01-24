DIST_DIR = dist

all: build

deps:
	go get "github.com/goamz/goamz/aws"

build: clean deps
	test -d $(DIST_DIR) || mkdir $(DIST_DIR)
	go build -o $(DIST_DIR)/r53dyndns main.go

clean:
	rm -rf $(DIST_DIR)

test:
	@go test -run=. -test.v ./config
	@go test -run=. -test.v ./utils
