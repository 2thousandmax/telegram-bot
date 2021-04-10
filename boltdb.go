package main

import (
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

func initBoltDB() (*bolt.DB, error) {
	db, err := bolt.Open(".bolt/bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Batch(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(UsersBucket))
		if err != nil {
			return err
		}

		return err
	}); err != nil {
		return nil, err
	}

	return db, nil
}
