run:
	go run main.go

clean:
	go clean && rm -rf bin/

build:
	go build -o bin/todo/todo main.go

compile:
	GOOS=linux GOARCH=386 go build -o bin/todo-linux-386/todo main.go
	GOOS=linux GOARCH=arm go build -o bin/todo-linux-arm/todo main.go
	GOOS=linux GOARCH=arm64 go build -o bin/todo-linux-arm64/todo main.go
	GOOS=linux GOARCH=amd64 go build -o bin/todo-linux-amd64/todo main.go
	GOOS=windows GOARCH=386 go build -o bin/todo-windows-386/todo main.go
	GOOS=windows GOARCH=arm go build -o bin/todo-windows-arm/todo main.go
	GOOS=windows GOARCH=arm64 go build -o bin/todo-windows-arm64/todo main.go
	GOOS=windows GOARCH=amd64 go build -o bin/todo-windows-amd64/todo main.go