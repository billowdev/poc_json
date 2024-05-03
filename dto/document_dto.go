package dto

import "poc_json/models"

type SDocumentVersionResponse struct {
	ID          uint         `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Version     uint         `json:"version"`
	VersionType string       `json:"version_type"`
	Value       models.JSONB `json:"value"`
}
