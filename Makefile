build:
	go mod download && CGO_ENABLED=0 GOOS=darwin go build -o ./.bin/app ./cmd/api/main.go

run: build
	docker-compose up --build server
