run:
	go run main.go

build:
	go build -o bin/todo main.go

compile:
	echo "Compiling..."
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go
	GOOS=windows GOARCH=arm go build -o bin/main-windows-arm main.go
	GOOS=windows GOARCH=arm64 go build -o bin/main-windows-arm64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 main.go