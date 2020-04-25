package rdb

import (
	"log"

	"github.com/jinzhu/gorm"
)

func Rollback(tx *gorm.DB) {
	if err := tx.Rollback().Error; err != nil {
		log.Println("failed to rollback: " + err.Error())
	}
}
