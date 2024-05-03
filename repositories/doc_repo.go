package repositories

import "gorm.io/gorm"

type (
	IDocumentRepoInfs interface {
		GetTest() string
	}
	repositoryDeps struct {
		db *gorm.DB
	}
)
func NewDocumentRepo(db *gorm.DB) IDocumentRepoInfs {
	return &repositoryDeps{
		db: db,
	}
}

func (r *repositoryDeps) GetTest() string {
	return "test"
}