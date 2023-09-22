package config


import (
	"os"
)


func SetEnvironmentConfg() {
	// Database config
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_SCHEMA", "search")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "root")

	// ElasticSearch config
	os.Setenv("ELASTIC_URL", "http://localhost:9200")
	os.Setenv("ELASTIC_INDEX", "search")
}