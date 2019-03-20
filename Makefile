.PHONY: test fmt compile-mac compile-linux build clean

test:
	go test

fmt:
	go fmt

compile-mac:
	GOOS=darwin go build -o build/mac/spm-junit-formatter main.go
compile-linux:
	GOOS=linux go build -o build/linux/spm-junit-formatter main.go

build: test compile-mac compile-linux

clean:
	go clean
	rm -rf build/
