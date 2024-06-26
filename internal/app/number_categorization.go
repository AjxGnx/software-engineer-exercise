package app

import (
	"fmt"

	"github.com/AjxGnx/software-engineer-exercise/internal/domain/dto"
	"github.com/AjxGnx/software-engineer-exercise/internal/domain/entity"
	"github.com/AjxGnx/software-engineer-exercise/internal/infra/adapters/pg/repository"
)

type NumberCategorization interface {
	Categorize(number int64) (entity.Categorization, error)
	GetByNumber(number int64) (entity.Categorization, error)
	Get(paginate dto.Paginate) (*entity.Paginator, error)
}

type numberCategorizationApp struct {
	repo repository.NumberCategorization
}

func NewNumberCategorizationApp(repo repository.NumberCategorization) NumberCategorization {
	return numberCategorizationApp{
		repo,
	}
}

func (app numberCategorizationApp) Categorize(number int64) (entity.Categorization, error) {
	Conditions := map[bool]string{
		number%3 == 0 && number%5 != 0: "type-1",
		number%3 != 0 && number%5 == 0: "type-2",
		number%3 == 0 && number%5 == 0: "type-3",
		number%3 != 0 && number%5 != 0: fmt.Sprintf("%v", number),
	}

	numberToCategorize := entity.Categorization{
		Number:   number,
		Category: Conditions[true],
	}

	categorization, err := app.repo.Insert(numberToCategorize)
	if err != nil {
		return categorization, err
	}

	return categorization, nil
}

func (app numberCategorizationApp) GetByNumber(number int64) (entity.Categorization, error) {
	return app.repo.GetByNumber(number)
}

func (app numberCategorizationApp) Get(paginate dto.Paginate) (*entity.Paginator, error) {
	return app.repo.Get(paginate.Page, paginate.Limit)
}
