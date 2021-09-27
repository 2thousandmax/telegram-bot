package main

import "database/sql"

type User struct {
	Id          int `json:"id"`
	TelegramId  int `json:"telegram_id"`
	TelegramTag int `json:"telegram_tag"`
	StudentId   int `json:"student_id"`
}

type StorageProvider interface {
	Users() StorageUserProvider
	Schedule() StorageScheduleProvider
}

type StorageUserProvider interface{
	Save(User) (int, error)
	Get(id int) (User, error)
}

type Storage struct {
	// Schedule ScheduleStorageProvider
}

func NewStorage(*sql.DB) *Storage {
	return &Storage{}
}

