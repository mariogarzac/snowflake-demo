package connection

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"log"
)

var conn Connection

type Connection struct {
	user         string
	password     string
	organization string
	account      string
	database     string
	schema       string
	warehouse    string
}

func (c *Connection) LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Could not load .env file")
	}

	conn.user = os.Getenv("USER")
	conn.password = os.Getenv("PASSWORD")
	conn.organization = os.Getenv("ORGANIZATION")
	conn.account = os.Getenv("ACCOUNT")
	conn.database = os.Getenv("DATABASE")
	conn.schema = os.Getenv("SCHEMA")
	conn.warehouse = os.Getenv("WAREHOUSE")
}

func (c *Connection) GetConnectionString() string {

	return fmt.Sprintf("%s:%s@%s-%s/%s/%s?warehouse=%s",
		conn.user,
		conn.password,
		conn.organization,
		conn.account,
		conn.database,
		conn.schema,
		conn.warehouse)
}
