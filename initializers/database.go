package initializers

import (
	"fmt"
	"log"

	"github.com/buabaj/post-api/models"
	"github.com/gobuffalo/envy"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	var DB_URL = envy.Get("DB_URL", "")
	dsn := DB_URL
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
