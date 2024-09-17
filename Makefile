run: build
	./bin/osmosis

build:
	go build -o bin/osmosis ./main.go

test:
	go test ./*.go
