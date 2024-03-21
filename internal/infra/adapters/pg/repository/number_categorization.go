package repository

import (
	"database/sql"

	"github.com/AjxGnx/software-engineer-exercise/internal/domain/entity"
	"github.com/AjxGnx/software-engineer-exercise/internal/infra/adapters/pg/queries"
	"github.com/labstack/gommon/log"
)

type NumberCategorization interface {
	Insert(categorization entity.Categorization) (entity.Categorization, error)
	GetByNumber(number int) error
	Get() error
}

type numberCategorizationRepo struct {
	db *sql.DB
}

func NewNumberCategorizationRepo(db *sql.DB) NumberCategorization {
	return numberCategorizationRepo{
		db,
	}
}

func (repo numberCategorizationRepo) Insert(categorization entity.Categorization) (entity.Categorization, error) {
	var rowInserted entity.Categorization

	err := repo.db.QueryRow(queries.InsertCategorization, categorization.Number, categorization.Category).
		Scan(&rowInserted.ID, &rowInserted.Number, &rowInserted.Category)
	if err != nil {
		log.Errorf("something happened inserting rows in categorization table %s", err)
		return rowInserted, err
	}

	return rowInserted, nil
}

func (repo numberCategorizationRepo) GetByNumber(number int) error {
	log.Info("APP - GetByNumber - ", number)
	return nil
}

func (repo numberCategorizationRepo) Get() error {
	log.Info("APP- GET")
	return nil
}