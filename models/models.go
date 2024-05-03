package models

import (
	"time"

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
	ID        uint            `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Title     string          `json:"title"`
	Subtitle  string          `json:"subtitle"`
	Message   string          `json:"message"`
	Status    ST_NOTIFICATION `json:"status"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
}

var TNNotification = "notifications"

func (st *NotificationModel) TableName() string {
	return TNNotification
}

const (
	EVT_NOTIFICATION = "notification"
	// EVT_DOCUMENT_CREATED = "document_created"
)

var Topics = []string{
	EVT_NOTIFICATION,
}

type NOTI_TYPE string

const (
	NOTI_TYPE_INFO    NOTI_TYPE = "info"
	NOTI_TYPE_WARNING NOTI_TYPE = "warning"
	NOTI_TYPE_ERROR   NOTI_TYPE = "error"
	NOTI_TYPE_SUCCESS NOTI_TYPE = "success"
)

type NOTI_STATUS string

const (
	NOTI_STATUS_UNREAD   NOTI_STATUS = "UNREAD"
	NOTI_STATUS_READ     NOTI_STATUS = "READ"
	NOTI_STATUS_ARCHIVED NOTI_STATUS = "ARCHIVED"
)

type KafkaEvent struct {
	EventType        string      `json:"event_type"`
	NotificationType NOTI_TYPE   `json:"notification_type"`
	Title            string      `json:"title"`
	Message          string      `json:"message"`
	Timestamp        time.Time   `json:"timestamp"`
	Status           NOTI_STATUS `json:"status"`
}

type STDocumentVersion struct {
	PortOfLoading     string `json:"port_of_loading"`
	PortOfDestination string `json:"port_of_destination"`
	CompanyName       string `json:"company_name"`
	Address           string `json:"address"`
}
