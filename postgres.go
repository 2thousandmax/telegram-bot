package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	DBPingTimeout     int = 5
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

func NewPostgresDatabase(p PostgresDatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", p.toDataSourceName())
	if err != nil {
		return &sql.DB{}, err
	}

	attempt := DBPingMaxAttempts
	for {
		attempt--
		err := db.Ping()
		if err != nil {
			log.Fatalln(err, "RETRIES LEFT:", attempt)

			if attempt <= 0 {
				time.Sleep( time.Duration(DBPingTimeout)* time.Second)

				continue
			}

			return nil, err
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
