package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"log"
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

	// Инициализация репозитория
	productRepository := repository.NewProductRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)

	// Инициализация use case
	productService := usecase.NewProductService(productRepository, categoryRepository)

	// Инициализация обработчика
	productHandler := handler.NewProductHandler(productService)

	// Инициализация Echo
	e := echo.New()

	productHandler.Register(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
