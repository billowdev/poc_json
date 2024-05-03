package repositories

import (
	"poc_json/models"

	"gorm.io/gorm"
)

type (
	INotiRepoInfs interface {
		GetAllNotifications() ([]models.NotificationModel, error)
		CreateNotification(payload *models.NotificationModel) error
	}
	notiRepoDeps struct {
		db *gorm.DB
	}
)

func NewNotiRepo(db *gorm.DB) INotiRepoInfs {
	return &notiRepoDeps{db: db}
}

// CreateNotification implements INotificationRepoInfs.
func (n *notiRepoDeps) CreateNotification(payload *models.NotificationModel) error {
	if err := n.db.Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// GetAllNotifications implements INotificationRepoInfs.
func (n *notiRepoDeps) GetAllNotifications() ([]models.NotificationModel, error) {
	panic("unimplemented")
}
