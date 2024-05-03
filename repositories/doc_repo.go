package repositories

import (
	"poc_json/models"

	"gorm.io/gorm"
)

type (
	IDocumentRepoInfs interface {
		GetTest() error

		HelperCreateDocumentVersion(tx *gorm.DB, p *models.DocumentVersionModel) error
		GetDocumentVersion(versionID string) (*models.DocumentVersionModel, error)
		BeginTransaction() (*gorm.DB, error)
	}
	repoDeps struct {
		db *gorm.DB
	}
)


func NewDocumentRepo(db *gorm.DB) IDocumentRepoInfs {
	return &repoDeps{
		db: db,
	}
}
// GetDocumentVersion implements IDocumentRepoInfs.
func (r *repoDeps) GetDocumentVersion(versionID string) (*models.DocumentVersionModel, error) {
	panic("unimplemented")
}

// HelperCreateDocumentVersion implements IDocumentRepoInfs.
func (r *repoDeps) HelperCreateDocumentVersion(tx *gorm.DB, p *models.DocumentVersionModel) error {
	if err := tx.Create(&p).Error; err != nil {
		return err
	}
	return nil
}

// BeginTransaction implements IDocumentRepoInfs.
func (r *repoDeps) BeginTransaction() (*gorm.DB, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (r *repoDeps) GetTest() error {
	return nil
}
