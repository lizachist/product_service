package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"product_service/internal/domain"
	"product_service/internal/dto"
	"strconv"
)

type ProductHandler struct {
	productService  domain.ProductService
	categoryService domain.CategoryService
	logger          *slog.Logger
}

func NewProductHandler(productUseCase domain.ProductService, categoryService domain.CategoryService, logger *slog.Logger) *ProductHandler {
	return &ProductHandler{
		productService:  productUseCase,
		categoryService: categoryService,
		logger:          logger,
	}
}

func (h *ProductHandler) Register(e *echo.Echo) {
	h.logger.Info("Registering product routes")
	e.POST("/products", h.CreateProduct)
	e.PUT("/products/:id", h.UpdateProduct)
	e.DELETE("/products/:id", h.DeleteProduct)
	e.GET("/products", h.GetAllProducts)
	h.logger.Info("Product routes registered successfully")
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	h.logger.Info("Handling CreateProduct request")
	user := new(domain.Product)
	if err := c.Bind(user); err != nil {
		h.logger.Error("Error binding request payload", "error", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.productService.Create(user); err != nil {
		h.logger.Error("Error creating product", "error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create product"})
	}

	h.logger.Info("Product created successfully", "id", user.ID)
	return c.JSON(http.StatusCreated, user)
}
func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	h.logger.Info("Handling UpdateProduct request")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid product ID", "error", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	product := new(domain.Product)
	if err := c.Bind(product); err != nil {
		h.logger.Error("Error binding request payload", "error", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	product.ID = id

	if err := h.productService.Update(product); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			h.logger.Warn("Product not found", "id", id)
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		h.logger.Error("Failed to update product", "error", err, "id", id)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update product"})
	}

	h.logger.Info("Product updated successfully", "id", id)
	return c.JSON(http.StatusOK, product)
}
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	h.logger.Info("Handling DeleteProduct request")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid product ID", "error", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	product := new(domain.Product)
	if err := c.Bind(product); err != nil {
		h.logger.Error("Error binding request payload", "error", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	product.ID = id

	if err := h.productService.Delete(product); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			h.logger.Warn("Product not found", "id", id)
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		h.logger.Error("Failed to delete product", "error", err, "id", id)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete product"})
	}

	h.logger.Info("Product deleted successfully", "id", id)
	return c.NoContent(http.StatusNoContent)
}

func (h *ProductHandler) GetAllProducts(c echo.Context) error {
	h.logger.Info("Handling GetAllProducts request")
	products, err := h.productService.GetAll()
	if err != nil {
		h.logger.Error("Failed to fetch products", "error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch products"})
	}

	// Создаем слайс для хранения DTO продуктов
	productResponses := make([]*dto.ProductResponse, 0, len(products))

	// Преобразуем каждый продукт в DTO
	for _, product := range products {
		// Получаем категорию для продукта
		category, err := h.categoryService.GetByID(*product.CategoryID)
		if err != nil {
			// Обрабатываем ошибку, если не удалось получить категорию
			h.logger.Error("Failed to fetch category", "error", err, "categoryID", *product.CategoryID)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch category"})
		}

		// Создаем DTO для продукта
		productResponse := dto.NewProductResponse(product, category)
		productResponses = append(productResponses, productResponse)
	}

	h.logger.Info("Successfully fetched all products", "count", len(productResponses))
	return c.JSON(http.StatusOK, productResponses)
}
