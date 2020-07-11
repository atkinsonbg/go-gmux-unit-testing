package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	backoff "github.com/cenkalti/backoff/v4"
	"os"
)

var db *sql.DB

func InitDB() {
	config := dbConfig()
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable",
		config["dbhost"], config["dbuser"], config["dbname"])
	log.Println(psqlInfo)
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

func dbConfig() map[string]string {
	conf := make(map[string]string)
	ok := false
	conf["dbhost"], ok = os.LookupEnv("DBHOST")
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	conf["dbname"], ok = os.LookupEnv("DBNAME")
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf["dbuser"], ok = os.LookupEnv("DBUSER")
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	return conf
}
