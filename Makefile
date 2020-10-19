run:
	go run main.go

build:
	go build -o bin/boiler .

install:
	./install.sh

test:
	go test -tags integration ./...
