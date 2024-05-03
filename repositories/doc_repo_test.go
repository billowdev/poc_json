package repositories_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"poc_json/models"
	"poc_json/repositories"
	"poc_json/utils"
)

func TestHelperCreateDocumentVersion(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	require.NoError(t, err)

	repo := repositories.NewDocumentRepo(gormDB)
	p, err := utils.NewJSONB(models.STDocumentVersion{
		PortOfLoading:     "HCM",
		PortOfDestination: "HAN",
		CompanyName:       "TEST",
		Address:           "test",
	})
	if err != nil {
		require.NoError(t, err)
	}
	docVersion := models.DocumentVersionModel{
		Version:     1,
		VersionType: "DRAFT",
		Value:       p,
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(
			`INSERT INTO "document_versions" ("created_at","updated_at","deleted_at","version","version_type","value") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id","id"`,
		)).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 1, "DRAFT", p).WillReturnRows(sqlmock.NewRows([]string{"id", "id"}).AddRow(1, 1))
		mock.ExpectCommit()

		tx := gormDB.Begin()
		err := repo.HelperCreateDocumentVersion(tx, &docVersion)
		t.Log(docVersion.CreatedAt)
		t.Log("Success case passed")
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO").WillReturnError(errors.New("error"))
		mock.ExpectRollback()

		tx := gormDB.Begin()
		err := repo.HelperCreateDocumentVersion(tx, &docVersion)
		require.Error(t, err)
	})
}
