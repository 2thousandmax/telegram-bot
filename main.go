package main

import (
	"log"
)

func main() {
	config, err := NewConfig("configs/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := NewPostgresDatabase(PostgresDatabaseConfig{
		host:     "",
		port:     0,
		user:     "",
		password: "",
		dbName:   "",
		sslMode:  "",
	})
	if err != nil {
		log.Fatalln(err)
	}

	storage := NewStorage(db)

	bot := NewTelegramBot(config, storage)

	if err := bot.Start(); err != nil {
		log.Fatalf("bot.Start: %v", err)
	}
}
