.SILENT:

run:
	go run src/cmd/api/main.go

build:
	docker build --tag origin-take-home .

run-docker:
	docker run --publish 8080:8080 origin-take-home
