package main

import (
	"gorm.io/gorm"
)



type DocumentCategoryModel struct {
	gorm.Model
	ID           uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	DocumentName string `json:"document_name"`
	Template     JSONB  `json:"template"`
}
var TNDocumentCategory = "document_categories"
func (st *DocumentCategoryModel) TableName() string {
	return TNDocumentCategory
}

type DocumentModel struct {
	gorm.Model
	ID                 uint                  `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	DocumentCategoryID string                `json:"document_category_id"`
	DocumentCategory   DocumentCategoryModel `gorm:"foreignkey:DocumentCategoryID"`
}

var TNDocument = "documents"

func (st *DocumentModel) TableName() string {
	return TNDocument
}

type DocumentVersionModel struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Version     uint   `json:"version"`
	VersionType string `json:"version_type"`
	Value       JSONB  `json:"value"`
}

var TNDocumentVersion = "document_versions"

func (st *DocumentVersionModel) TableName() string {
	return TNDocumentVersion
}

type ST_NOTIFICATION string

const (
	NS_NEW     ST_NOTIFICATION = "NEW"
	NS_INFO    ST_NOTIFICATION = "INFO"
	NS_ERROR   ST_NOTIFICATION = "ERROR"
	NS_WARNING ST_NOTIFICATION = "WARNING"
)

type NotificationModel struct {
	gorm.Model
	ID       uint            `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Title    uint            `json:"title"`
	Subtitle string          `json:"subtitle"`
	Status   ST_NOTIFICATION `json:"status" gorm:"type:ENUM('new', 'info', 'error', 'warning')"`
}

var TNNotification = "notifications"

func (st *NotificationModel) TableName() string {
	return TNNotification
}
