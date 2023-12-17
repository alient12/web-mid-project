package basketrepo

import (
	"context"
	"errors"
	"myapp/internal/domain/model"
)

var ErrBasketIDDuplicate = errors.New("basket id already exists")
var ErrBasketStateCompleted = errors.New("basket already completed")

type GetCommand struct {
	ID *uint64
}

type Repository interface {
	Add(ctx context.Context, model model.Basket) error
	Get(ctx context.Context, cmd GetCommand) []model.Basket
	GetAll(ctx context.Context) []model.Basket
	Update(ctx context.Context, model model.Basket) error
	Delete(ctx context.Context, id uint64) error
}
