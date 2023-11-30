default: test

build:
	go build -buildmode=plugin -o plugin.so ./so/main.go
	go build -o plugin.cli cli/main.go

test: build
	go test -bench . -benchmem
