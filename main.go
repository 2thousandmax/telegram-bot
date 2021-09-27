package main

import (
	"log"
)

func main() {
	config, err := NewConfig("configs/config.yaml")
	if err != nil {
		log.Panic(err)
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
		log.Fatalf("Database error: %v", err)
	}

	storage := NewStorage(db)

	bot := NewTelegramBot(config.Bot, config.Replies, storage)

	if err := bot.Start(); err != nil {
		log.Fatalf("Bot error: %v", err)
	}
}
