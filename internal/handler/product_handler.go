package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"product_service/internal/domain"
	"strconv"
)

type ProductHandler struct {
	productService domain.ProductService
}

func NewProductHandler(productUseCase domain.ProductService) *ProductHandler {
	return &ProductHandler{productService: productUseCase}
}

func (h *ProductHandler) Register(e *echo.Echo) {
	e.POST("/products", h.CreateProduct)
	e.PUT("/products/:id", h.UpdateProduct)
	e.DELETE("/products/:id", h.DeleteProduct)
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	user := new(domain.Product)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.productService.Create(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create product"})
	}

	return c.JSON(http.StatusCreated, user)
}
func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	product := new(domain.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	product.ID = id

	if err := h.productService.Update(product); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update product"})
	}

	return c.JSON(http.StatusOK, product)
}
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	product := new(domain.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	product.ID = id

	if err := h.productService.Delete(product); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete product"})
	}

	return c.NoContent(http.StatusNoContent)
}
