package app

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"heloo-go/internal/handler"
	"heloo-go/internal/repository"
	"heloo-go/internal/service"
)

func NewServer() (*echo.Echo, *gorm.DB, error) {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, nil, echo.NewHTTPError(http.StatusInternalServerError, "DATABASE_URL is required")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	itemRepo := repository.NewItemRepository(db)
	if err := itemRepo.AutoMigrate(); err != nil {
		return nil, nil, err
	}
	itemSvc := service.NewItemService(itemRepo)
	itemHandler := handler.NewItemHandler(itemSvc)

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	// Prefixed routes for Netlify (/api/*)
	api := e.Group("/api")
	api.GET("/items", itemHandler.List)
	api.GET("/items/:id", itemHandler.Get)
	api.POST("/items", itemHandler.Create)
	api.PUT("/items", itemHandler.Update)
	api.DELETE("/items", itemHandler.Delete)

	log.Println("server initialized")
	return e, db, nil
}
