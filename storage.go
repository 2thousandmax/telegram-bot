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

func (s *UserStorage) Save(id int, username string) error {
	if err := s.db.Batch(func(tx *bolt.Tx) error {
		u, err := tx.Bucket([]byte("users")).CreateBucketIfNotExists(itob(id))
		if err != nil {
			return err
		}

		if err := u.Put([]byte("username"), []byte(username)); err != nil {
			return err
		}

		if err := u.Put([]byte("group"), []byte("")); err != nil {
			return err
		}

		return err
	}); err != nil {
		return err
	}

	return nil
}

func (s *UserStorage) GetGroup(id int, bucket Bucket) (string, error) {
	var group string

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket)).Bucket(itob(id))
		group = string(b.Get([]byte("group")))
		return nil
	})

	if group == "" {
		err = errorInvalidGroup
	}

	return group, err
}

func (s *UserStorage) SetGroup(id int, group string, bucket Bucket) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket)).Bucket(itob(id))

		if err := b.Put([]byte("group"), []byte(group)); err != nil {
			return err
		}

		return nil
	})

	return err
}
