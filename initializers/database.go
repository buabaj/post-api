package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/buabaj/post-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dbURL := os.Getenv("DB_URL")
	dsn := dbURL
	fmt.Println(dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
}

func GetDB() *gorm.DB {
	return DB
}

func EmptyDB() {
	DB.Migrator().DropTable(&models.Post{})
}
