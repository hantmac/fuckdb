all: gotool
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v .
install:
	go install
clean:
	rm -f DbToGoStruct
	rm -f gin-bin
	rm -rf logs
gotool:
	gofmt -w .
	go vet . | grep -v vendor;true

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make install - install dependency"

.PHONY: clean gotool ca help