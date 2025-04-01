package config

import (
	"github.com/mariogarzac/snowflake_demo/internal/db"
)

type Config struct {
	Database *db.DB
}
