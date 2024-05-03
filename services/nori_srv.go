package services

import (
	"poc_json/models"
	"poc_json/repositories"
)

type (
	INotiSrvInfs interface {
		CreateNotification(payload *models.NotificationModel) error
	}
	notiSrvDeps struct {
		repo repositories.INotiRepoInfs
	}
)

func NewNotiSrv(repo repositories.INotiRepoInfs) INotiSrvInfs {
	return &notiSrvDeps{
		repo: repo,
	}
}

func (s *notiSrvDeps) CreateNotification(payload *models.NotificationModel) error {
	if err := s.repo.CreateNotification(payload); err != nil {
		return err
	}
	return nil
}
