run: build
	./bin/osmosis

build:
	go build -o bin/osmosis ./main.go

test:
	go test ./*.go

format:
	gofmt -s -w .

vet:
	go vet ./*.go

fix:
	go fix .

clean:
	rm -rf bin
