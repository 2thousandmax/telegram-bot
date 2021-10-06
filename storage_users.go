package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const userTable = "users"

type UserStorage struct {
	db *sqlx.DB
}

var _ UserStorageProvider = (*UserStorage)(nil)

func (us *UserStorage) Create(user User) (int, error) {
	query := `
	INSERT INTO ? 
	 (id, telegram_id, telegram_tag,)
	 VALUES(?,?,?)
	RETURNING id;`

	row:= us.db.QueryRow(query, userTable, user.Id, user.TelegramId, user.TelegramTag)
	
	var id int
	if err := row.Scan(&id); err != nil{
		return 0, fmt.Errorf("sql.Row.Scan: %v", err)
	}

	return id, nil
}

func (us *UserStorage) Get(telegramId int) (User, error) {
	user := User{}
	
	query := `
	SELECT * 
	FROM ?
	WHERE telegram_id=?`

	if err := us.db.Select(&user, query, userTable, telegramId); err != nil{
		return User{}, fmt.Errorf("db.Select: %v", err)
	}

	return user, nil
}

func NewUserStorage(db *sqlx.DB) UserStorageProvider {
	userStorage := UserStorage{
		db: db,
	}

	return &userStorage
}
