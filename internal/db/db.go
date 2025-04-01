package db

import (
	"database/sql"
	"fmt"

	"github.com/mariogarzac/snowflake_demo/pkg/connection"
	_ "github.com/snowflakedb/gosnowflake"

	"log"
)

type DB struct {
	Database *sql.DB
}

func NewConnection() *DB {
	return &DB{
		Database: connect(),
	}
}

func connect() *sql.DB {
	conn := connection.Connection{}

	conn.LoadEnvVariables()
	connectionString := conn.GetConnectionString()

	db, err := sql.Open("snowflake", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected successfully")

	defer db.Close()

	return db
}

func (db *DB) GetAllRecords() {
	rows, err := db.Database.Query("SELECT tweet_id, text FROM tweets LIMIT 10")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tweetID string
		var text string
		err = rows.Scan(&tweetID, &text)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tweetID, text)
	}
}
