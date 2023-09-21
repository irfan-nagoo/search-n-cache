package component

import (
	"os"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConnection *gorm.DB

func InitializeDBConnection() *gorm.DB {
	log.Info("Connecting to the Database")
	connectionString := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
		"@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_SCHEMA") +
		"?charset=utf8mb4&parseTime=True&loc=Local"

	dbConn, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Error("Error occured while connecting to the database")
		log.Panic(err)
		return nil
	}

	log.Info("Connected to database Successfully")
	return dbConn
}
