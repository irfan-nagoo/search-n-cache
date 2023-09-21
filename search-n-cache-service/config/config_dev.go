package config


import (
	"os"
)


func SetEnvironmentConfg() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_SCHEMA", "search")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "root")
}