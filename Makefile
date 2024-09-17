run: build
	./bin/osmosis

build:
	go build -o bin/osmosis ./main.go

test:
	go test ./*.go

lint:
	gofmt -s -w .

vet:
	go vet ./*.go

fix:
	go fix ./*.go

clean:
	rm -rf bin
