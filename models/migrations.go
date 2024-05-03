package models

import "gorm.io/gorm"

func RunMigrations(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
		if err != nil {
			return err
		}
		err = tx.AutoMigrate(
			// TODO START USER
			&DocumentCategoryModel{},
			&DocumentModel{},
			&DocumentVersionModel{},
			&NotificationModel{},
		)
		return err
	})

	return err
}
