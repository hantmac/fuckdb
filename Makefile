all: gotool
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fuckdb -v .
install:
	go install
clean:
	rm -f DbToGoStruct
	rm -f gin-bin
	rm -rf logs
gotool:
	gofmt -w .
	go vet . | grep -v vendor;true
build-frontend:
	cd frontend && npm run build
start:
	go run main.go

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make install - install dependency"
	@echo "make start - start backend service"
	@echo "make build-frontend - build frontend"

.PHONY: clean gotool ca help