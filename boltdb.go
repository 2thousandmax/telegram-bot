package main

import (
	"os"
	"time"

	"github.com/boltdb/bolt"
)

type Bucket string

const (
	UsersBucket Bucket = "users"
)

type User struct {
	ID        string
	UserName  string
	Group     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	path = "bolt"
	mode = 0700
)

func initBoltDB() (*bolt.DB, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
	}

	db, err := bolt.Open("bolt/bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}
