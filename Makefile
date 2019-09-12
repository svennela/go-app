all: linux darwin

darwin: dirs
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/my-go-app .

linux: dirs
	GOOS=linux GOARCH=amd64 go build -o bin/linux/my-go-app .

dirs:
	mkdir -p bin/linux
	mkdir -p bin/darwin
