run:
	go run main.go

clean:
	go clean && rm -rf bin/

build:
	go build -o bin/todo main.go