package repository

import (
	"database/sql"
	"math"

	"github.com/AjxGnx/software-engineer-exercise/internal/domain/entity"
	"github.com/AjxGnx/software-engineer-exercise/internal/infra/adapters/pg/queries"
	"github.com/labstack/gommon/log"
)

type NumberCategorization interface {
	Insert(categorization entity.Categorization) (entity.Categorization, error)
	GetByNumber(number int64) (entity.Categorization, error)
	Get(page int, limit int) (*entity.Paginator, error)
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

func (repo numberCategorizationRepo) GetByNumber(number int64) (entity.Categorization, error) {
	var categorization entity.Categorization

	err := repo.db.QueryRow(queries.GetByNumber, number).
		Scan(&categorization.ID, &categorization.Number, &categorization.Category)

	if err != nil {
		log.Printf("error fetching categorization by number: %s", err)
		return categorization, err
	}

	return categorization, nil
}

func (repo numberCategorizationRepo) Get(page int, limit int) (*entity.Paginator, error) {
	var categorizations []entity.Categorization

	offset := (page - 1) * limit

	rows, err := repo.db.Query(queries.GetWithLimitAndOffset, limit, offset)
	if err != nil {
		log.Printf("error fetching categorizations for page %d: %s", page, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var categorization entity.Categorization
		if err := rows.Scan(&categorization.ID, &categorization.Number, &categorization.Category); err != nil {
			log.Printf("error scanning categorization row: %s", err)
			return nil, err
		}
		categorizations = append(categorizations, categorization)
	}

	totalRecords, err := repo.countTotalRecords()
	if err != nil {
		log.Printf("error getting total count of categorizations: %s", err)
		return nil, err
	}

	paginator := &entity.Paginator{
		TotalRecord: totalRecords,
		TotalPage:   int(math.Ceil(float64(totalRecords) / float64(limit))),
		Records:     categorizations,
		Offset:      offset,
		Limit:       limit,
		Page:        page,
	}

	if page > 1 {
		paginator.PrevPage = page - 1
	} else {
		paginator.PrevPage = page
	}

	if page == paginator.TotalPage {
		paginator.NextPage = page
	} else {
		paginator.NextPage = page + 1
	}

	return paginator, nil
}

func (repo numberCategorizationRepo) countTotalRecords() (int, error) {
	var total int

	if err := repo.db.QueryRow(queries.GetCount).Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
