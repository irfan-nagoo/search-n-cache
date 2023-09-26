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
	os.Setenv("ELASTIC_URL", "https://localhost:9200")
	os.Setenv("ELASTIC_SKIP_VERIFY_SSL", "true")
	os.Setenv("ELASTIC_USERNAME", "elastic")
	os.Setenv("ELASTIC_PASSWORD", "elastic")
	os.Setenv("ELASTIC_INDEX", "search")

	// Redis config
	os.Setenv("REDIS_URL", "localhost:6379")
	os.Setenv("REDIS_SKIP_VERIFY_SSL", "true")
	os.Setenv("REDIS_CACHE_EXPIRY_INTERVAL_SEC", "120")
	os.Setenv("REDIS_USERNAME", "")
	os.Setenv("REDIS_PASSWORD", "")
	os.Setenv("REDIS_DB", "0")

}
