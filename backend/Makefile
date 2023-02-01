BINARY_NAME=myapp

install:
	go get -u && go install

format:
	gofmt -w .

lint:
	golint -set_exit_status .

build:
	go build -o bin/$(BINARY_NAME)

clean:
	rm -f bin/$(BINARY_NAME)

run:
	./bin/$(BINARY_NAME)

test:
	go test -v ./...