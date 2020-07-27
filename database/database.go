package database

import (
	"database/sql"
	"fmt"
	backoff "github.com/cenkalti/backoff/v4"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

type Config struct {
	User        string `envconfig:"DBUSER"`
	Name        string `envconfig:"DBNAME"`
	Host     	string `envconfig:"DBHOST"`
}

func InitDB() {
	c := dbConfig()
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable",
		c.Host, c.User, c.Name)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}

	err = backoff.Retry(pingDb, backoff.NewExponentialBackOff())
	if err != nil {
		log.Panic(err)
	}
}

var pingDb backoff.Operation = func() error {
	err := db.Ping()
	if err != nil {
		log.Println("DB is not ready...backing off...")
		return err
	}
	log.Println("DB is ready!")
	return nil
}

func dbConfig() Config {
	var c Config
	err := envconfig.Process("myapp", &c)
	if err != nil {
		log.Panic(err)
	}
	return c
}
