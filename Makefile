.DEFAULT_GOAL := run
dep:
	go mod download

build: dep
	go build -o ./.bin/telegram-bot ./main.go

run:
	go run main.go