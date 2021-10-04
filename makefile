.SILENT:

run:
	go run src/cmd/api/main.go

test:
	go test ./...

build:
	docker build --tag origin-take-home .

run-docker: build
	docker run --publish 8080:8080 origin-take-home
