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
	cd frontend && npm install && npm run build
build: build-frontend
	rm -rf cmd/fuckdb/cmd/dist
	cp -r frontend/dist cmd/fuckdb/cmd
	cd cmd/fuckdb && go build -o fuckdb
	mv cmd/fuckdb/fuckdb .
linux: build-frontend
	rm -rf cmd/fuckdb/cmd/dist
	cp -r frontend/dist cmd/fuckdb/cmd
	cd cmd/fuckdb && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fuckdb
	mv cmd/fuckdb/fuckdb .
darwin: build-frontend
	rm -rf cmd/fuckdb/cmd/dist
	cp -r frontend/dist cmd/fuckdb/cmd
	cd cmd/fuckdb && CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o fuckdb
	mv cmd/fuckdb/fuckdb .
win: build-frontend
	rm -rf cmd/fuckdb/cmd/dist
	cp -r frontend/dist cmd/fuckdb/cmd
	cd cmd/fuckdb && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o fuckdb
	mv cmd/fuckdb/fuckdb .
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