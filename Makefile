.DEFAULT_GOAL := run
dep:
	go mod download

build: dep
	go build -o ./.bin/telegram-bot.exe

run:
	go run .