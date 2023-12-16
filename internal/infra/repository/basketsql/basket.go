package basketsql

import (
	"context"
	"errors"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository/basketrepo"

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
