package main

import "github.com/jmoiron/sqlx"

type ID int

type User struct {
	Id          int `json:"id"`
	TelegramId  int `json:"telegram_id"`
	TelegramTag int `json:"telegram_tag"`
	StudentId   int `json:"student_id"`
}

type Student struct {
	Id      int `json:"id"`
	UserId  int `json:"user_id"`
	GroupId int `json:"group_id"`
}

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Subject struct {
	Id      int    `json:"id"`
	GroupId int    `json:"group_id"`
	Name    string `json:"name"`
}

type Class struct {
	Id        int      `json:"id"`
	GroupId   int      `json"group_id"`
	SubjectId int      `json:"subject_id"`
	Index     string   `json:"index"`
	Type      string   `json:"type"`
	Room      string   `json:"room"`
	Lecturer  Lecturer `json:"lecturer`
}

type Lecturer struct {
	Id         int    `json:"id"`
	Groups     []int  `json:"groups"`
	FirstName  string `json:"first_name"`
	SecondName string `json;"second_name"`
	LastName   string `json:"last_name"`
}

type Schedule struct {
	GroupId    int    `json:"group_id"`
	Weekday    string `json:"weekday"`
	WeekNumber string `json:"week_number`
	Classes    Class  `json:"classes"`
}

type Storage struct {
	Users    UserStorageProvider
	Schedule ScheduleStorageProvider
}

type UserStorageProvider interface {
	Create(User) (int, error)
	Get(id int) (User, error)
}

type ScheduleStorageProvider interface {
	Get(groupId int) (Schedule, error)
	Update()
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Users:    NewUserStorage(db),
		Schedule: nil,
	}
}
