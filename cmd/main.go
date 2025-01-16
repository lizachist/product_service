package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"log/slog"
	"os"
	"product_service/internal/handler"
	"product_service/internal/repository"
	"product_service/internal/usecase"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Подключение к базе данных
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Создаем логгер
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	slog.SetDefault(logger)

	// Инициализация репозитория
	productRepository := repository.NewProductRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)

	// Инициализация use case
	productService := usecase.NewProductService(productRepository, categoryRepository)
	categoryService := usecase.NewCategoryService(categoryRepository)

	// Инициализация обработчика
	productHandler := handler.NewProductHandler(productService, categoryService, logger)

	// Инициализация Echo
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	productHandler.Register(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
