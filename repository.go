package main

import "gorm.io/gorm"

type (
	IRepositoryInfs interface {
		GetTest() string
	}
	repositoryDeps struct {
		db *gorm.DB
	}
)
func NewRepository(db *gorm.DB) IRepositoryInfs {
	return &repositoryDeps{
		db: db,
	}
}

func (r *repositoryDeps) GetTest() string {
	return "test"
}