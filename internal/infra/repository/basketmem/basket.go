package basketmem

import (
	"context"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository/basketrepo"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type Repository struct {
	baskets map[uint64]model.Basket
	lock    sync.RWMutex
}

func New() *Repository {
	return &Repository{
		baskets: make(map[uint64]model.Basket),
		lock:    sync.RWMutex{},
	}
}

func (r *Repository) Add(_ context.Context, model model.Basket) error {
	r.lock.RLock()
	if _, ok := r.baskets[model.ID]; ok {
		return basketrepo.ErrBasketIDDuplicate
	}
	r.lock.RUnlock()

	r.lock.Lock()
	r.baskets[model.ID] = model
	r.lock.Unlock()

	return nil
}

func (r *Repository) Get(_ context.Context, cmd basketrepo.GetCommand) []model.Basket {
	r.lock.RLock()
	defer r.lock.RUnlock()

	var baskets []model.Basket

	if cmd.ID != nil {
		basket, ok := r.baskets[*cmd.ID]
		if !ok {
			return nil
		}
		baskets = []model.Basket{basket}
	} else {
		for _, basket := range r.baskets {
			baskets = append(baskets, basket)
		}
	}
	return baskets
}

func (r *Repository) GetAll(_ context.Context) []model.Basket {
	r.lock.RLock()
	defer r.lock.RUnlock()

	var baskets []model.Basket

	for _, basket := range r.baskets {
		baskets = append(baskets, basket)
	}

	return baskets
}

func (r *Repository) Update(_ context.Context, m model.Basket) error {
	r.lock.RLock()
	defer r.lock.RUnlock()

	if entry, ok := r.baskets[m.ID]; !ok {
		return echo.ErrNotFound
	} else {
		if entry.State {
			return basketrepo.ErrBasketStateCompleted
		}
		entry.Data = m.Data
		entry.State = m.State
		entry.UpdatedAt = time.Now()
	}

	return nil
}
