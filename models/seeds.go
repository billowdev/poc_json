package models


import (
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

func RunSeeds(db *gorm.DB) {
	SeedDocCat(db)

}

func SeedDocCat(db *gorm.DB) error {
	DOC_CAT := `[
	{
		"id": 1,
		"document_name": "Document 1",
		"template": {"key1": "value1"}
	},
	{
		"id": 2,
		"document_name": "Document 2",
		"template": {"key2": "value2"}
	},
	{
		"id": 3,
		"document_name": "Document 3",
		"template": {"key3": "value3"}
	},
	{
		"id": 4,
		"document_name": "Document 4",
		"template": {"key4": "value4"}
	},
	{
		"id": 5,
		"document_name": "Document 5",
		"template": {"key5": "value5"}
	}
]`
	var data []DocumentCategoryModel
	err := json.Unmarshal([]byte(DOC_CAT), &data)
	if err != nil {
		log.Fatal(err)
	}

	tx := db.Begin()

	// Loop through the countries and create records within the transaction
	for _, element := range data {
		var existing DocumentCategoryModel
		// Check if the country already exists in the database
		if err := tx.Where("id = ?", element.ID).First(&existing).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				// If the error is not 'Record Not Found', roll back the transaction and return the error
				tx.Rollback()
				return err
			}
			// If the country does not exist, create a new record
			if err := tx.Create(&element).Error; err != nil {
				// If there is an error in creating the record, roll back the transaction and return the error
				tx.Rollback()
				return err
			}
		}

		if existing.ID != 0 {
			if err := tx.Model(DocumentCategoryModel{}).
				Where("id = ?", existing.ID).
				Updates(DocumentCategoryModel{DocumentName: existing.DocumentName, Template: existing.Template}).
				Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil

}
