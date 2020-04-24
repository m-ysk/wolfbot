package initializer

import (
	"log"
	"time"
	"wolfbot/infra/rdb"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB(dbURL string) *gorm.DB {
	db := connectToDB(dbURL)
	migrate(db)

	return db
}

func connectToDB(dbURL string) *gorm.DB {
	var count = 0
	for {
		if count == 3 {
			log.Fatal("failed to connect to DB")
		}

		db, err := gorm.Open("postgres", dbURL)
		if err != nil {
			count++
			log.Println("retrying...")
			time.Sleep(3 * time.Second)
			continue
		}

		return db
	}
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(rdb.Village{})
}
