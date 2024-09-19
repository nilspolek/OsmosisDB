run: build
	./bin/osmosis

build:
	go build -o bin/osmosis ./main.go

test:
	go test ./paser
	go test ./database

format:
	gofmt -s -w .

vet:
	go vet ./*.go

fix:
	go fix .

clean:
	rm -rf bin
