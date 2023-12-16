package main

import (
	"log"
	"myapp/internal/domain/repository/basketrepo"
	"myapp/internal/infra/http/handler"
	"myapp/internal/infra/repository/basketsql"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("baskets.db"), new(gorm.Config))
	if err != nil {
		log.Fatalf("failed to connect database %v", err)
	}

	if err := db.AutoMigrate(new(basketsql.BasketDTO)); err != nil {
		log.Fatalf("failed to run migrations %v", err)
	}

	app := echo.New()

	// var repo basketrepo.Repository = basketmem.New()
	var repo basketrepo.Repository = basketsql.New(db)

	h := handler.NewBasket(repo)
	h.Register(app.Group("/v1"))

	if err := app.Start("127.0.0.1:1373"); err != nil {
		log.Fatalf("server failed to start %v", err)
	}
}
