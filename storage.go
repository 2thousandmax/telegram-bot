package main

import (
	"github.com/boltdb/bolt"
)

type UserStorage struct {
	db *bolt.DB
}

func NewUserStorage(db *bolt.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (s *UserStorage) SaveGroup(id int, group string) error {
	err := s.db.Batch(func(tx *bolt.Tx) error {
		u, err := tx.CreateBucketIfNotExists(itob(id))
		if err != nil {
			return err
		}

		return u.Put([]byte("group"), []byte(group))
	})

	return err
}

func (s *UserStorage) GetGroup(id int) (string, error) {
	var group string

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(itob(id))
		group = string(b.Get([]byte("group")))
		return nil
	})

	return group, err
}
