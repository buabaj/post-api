package initializers

import (
	"log"

	"github.com/buabaj/post-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	//dbURL := os.Getenv("DB_URL")
	dsn := "host=dpg-cejep9cgqg4ekmcq9hjg-a user=buabaj password=fJhNHrIlGQJJjqmI6s0k5LD4q4RUkzXw dbname=post_db_01n8 port=5432 sslmode=require"
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
