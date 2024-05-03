package utils

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func HandleTransaction(tx *gorm.DB, err error, featureName string) {
	if r := recover(); r != nil {
		msg := fmt.Sprintf("--------------%s-----------------", featureName)
		log.Println(msg)
		err := tx.Rollback()
		log.Println(err)
		log.Println("-------------------------------")
	} else if err != nil {
		tx.Rollback()
	} else {
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			log.Println(err)
		}
	}
}
