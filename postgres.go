package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	DBPingTimeout     int  = 5
	DBPingMaxAttempts uint = 5
)

type PostgresDatabaseConfig struct {
	host     string
	port     uint
	user     string
	password string
	dbName   string
	sslMode  string
}

func NewPostgresDatabase(p PostgresDatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", p.toDataSourceName())
	if err != nil {
		return &sqlx.DB{}, fmt.Errorf("sqlx.Open: %v", err)
	}

	attempt := DBPingMaxAttempts
	for {
		attempt--
		err := db.Ping()
		if err != nil {
			log.Println(err, "RETRIES LEFT:", attempt)

			if attempt <= 0 {
				time.Sleep(time.Duration(DBPingTimeout) * time.Second)

				continue
			}

			return nil, fmt.Errorf("NewPostgresDatabase: %v", err)
		}

		break
	}

	return db, nil
}

func (pc *PostgresDatabaseConfig) toDataSourceName() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pc.host,
		pc.port,
		pc.user,
		pc.password,
		pc.dbName,
		pc.sslMode,
	)
}
