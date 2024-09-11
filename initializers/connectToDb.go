package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
    var err error
    dsn := os.Getenv("DB_URL")
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // sin la declaración `var`

    if err != nil {
        panic("Failed to connect to database")
    }

    log.Print("Connected")
}
