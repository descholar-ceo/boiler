run:
	go run main.go

build:
	go build -o bin/boiler main.go

install:
	./install.sh

test:
	go test
