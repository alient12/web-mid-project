package basketsql

import (
	"context"
	"errors"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository/basketrepo"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BasketDTO struct {
	model.Basket
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Add(ctx context.Context, model model.Basket) error {
	tx := r.db.WithContext(ctx).Create(&BasketDTO{Basket: model})
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrDuplicatedKey) {
			return basketrepo.ErrBasketIDDuplicate
		}
		return tx.Error
	}
	return nil
}

func (r *Repository) Get(_ context.Context, cmd basketrepo.GetCommand) []model.Basket {
	var basketDTOs []BasketDTO
	var condition BasketDTO
	if cmd.ID != nil {
		condition.ID = *cmd.ID
	}
	if err := r.db.Where(&condition).Find(&basketDTOs); err != nil {
		// return nil
	}
	baskets := make([]model.Basket, len(basketDTOs))

	for index, dto := range basketDTOs {
		baskets[index] = dto.Basket
	}
	return baskets
}

func (r *Repository) GetAll(_ context.Context) []model.Basket {
	var basketDTOs []BasketDTO
	if err := r.db.Find(&basketDTOs); err != nil {
		// return nil
	}
	baskets := make([]model.Basket, len(basketDTOs))

	for index, dto := range basketDTOs {
		baskets[index] = dto.Basket
	}
	return baskets
}

func (r *Repository) Update(ctx context.Context, model model.Basket) error {
	var basketDTOs []BasketDTO
	var condition BasketDTO

	condition.ID = model.ID
	if err := r.db.Where(&condition).Find(&basketDTOs); err != nil {
		// return echo.ErrBadRequest
	}
	if len(basketDTOs) == 0 {
		return echo.ErrInternalServerError
	}
	if len(basketDTOs) > 1 {
		return echo.ErrInternalServerError
	}
	if basketDTOs[0].State {
		return basketrepo.ErrBasketStateCompleted
	}
	tx := r.db.WithContext(ctx).Model(&basketDTOs[0]).Update("Data", model.Data).Update("State", model.State).Update("UpdatedAt", time.Now())
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id uint64) error {
	var basketDTOs []BasketDTO
	var condition BasketDTO

	condition.ID = id
	if err := r.db.Where(&condition).Find(&basketDTOs); err != nil {
		// return echo.ErrBadRequest
	}
	if len(basketDTOs) == 0 {
		return echo.ErrInternalServerError
	}
	if len(basketDTOs) > 1 {
		return echo.ErrInternalServerError
	}

	tx := r.db.WithContext(ctx).Delete(&basketDTOs[0])
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
