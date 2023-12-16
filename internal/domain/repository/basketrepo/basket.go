package basketrepo

import (
	"context"
	"errors"
	"myapp/internal/domain/model"
)

var ErrBasketIDDuplicate = errors.New("basket id already exists")

type GetCommand struct {
	ID *uint64
}

type Repository interface {
	Add(ctx context.Context, model model.Basket) error
	Get(ctx context.Context, cmd GetCommand) []model.Basket
}
