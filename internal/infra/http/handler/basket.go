package handler

import (
	"errors"
	"log"
	"math/rand"
	"myapp/internal/domain/model"
	"myapp/internal/domain/repository/basketrepo"
	"myapp/internal/infra/http/request"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Basket struct {
	repo basketrepo.Repository
}

func NewBasket(repo basketrepo.Repository) *Basket {
	return &Basket{
		repo: repo,
	}
}

func (b *Basket) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	baskets := b.repo.Get(c.Request().Context(), basketrepo.GetCommand{
		ID: &id,
	})
	if len(baskets) == 0 {
		return echo.ErrNotFound
	}

	if len(baskets) > 1 {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, baskets[0])
}

func (b *Basket) Create(c echo.Context) error {
	var req request.BasketCreate

	if err := c.Bind(&req); err != nil {
		log.Print("cannot bind")
		return echo.ErrBadRequest
	}
	if err := req.Validate(); err != nil {
		log.Print("cannot validate")
		return echo.ErrBadRequest
	}

	id := rand.Uint64() % 1_000_000

	if err := b.repo.Add(c.Request().Context(), model.Basket{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Data:      req.Data,
		State:     req.State,
	}); err != nil {
		if errors.Is(err, basketrepo.ErrBasketIDDuplicate) {
			log.Print("duplicate id")
			return echo.ErrBadRequest
		}
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusCreated, id)
}

func (b *Basket) Get(c echo.Context) error {
	baskets := b.repo.GetAll(c.Request().Context())
	if len(baskets) == 0 {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, baskets)
}

func (b *Basket) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	var req request.BasketCreate

	if err := c.Bind(&req); err != nil {
		log.Print("cannot bind")
		return echo.ErrBadRequest
	}
	if err := req.Validate(); err != nil {
		log.Print("cannot validate")
		return echo.ErrBadRequest
	}

	if err := b.repo.Update(c.Request().Context(), model.Basket{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Data:      req.Data,
		State:     req.State,
	}); err != nil {
		if errors.Is(err, basketrepo.ErrBasketStateCompleted) {
			return echo.ErrBadRequest
		}
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusCreated, id)
}

func (b *Basket) Register(g *echo.Group) {
	g.POST("/basket", b.Create)
	g.GET("/basket", b.Get)
	g.GET("/basket/:id", b.GetByID)
	g.PATCH("/basket/:id", b.Update)
}
