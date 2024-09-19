run: build
	./bin/osmosis

build:
	go build -o bin/osmosis ./main.go

test:
	go test ./paser
	go test ./database

lint:
	go install golang.org/x/lint/golint@latest
	golint -set_exit_status ./...


vet:
	go vet ./*.go

fix:
	go fix .

clean:
	rm -rf bin
