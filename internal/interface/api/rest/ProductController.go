package rest

import (
	"net/http"

	"github.com/AureoFJunior/muguiwaraGo/internal/application/interfaces"
	"github.com/AureoFJunior/muguiwaraGo/internal/interface/api/rest/request"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	service interfaces.ProductService
}

func NewProductController(e *echo.Echo, service interfaces.ProductService) *ProductController {
	controller := &ProductController{
		service: service,
	}

	e.POST("/products", controller.CreateProduct)
	e.GET("/products", controller.GetAllProducts)
	e.GET("/products/:id", controller.GetProductByID)

	return controller
}

func (pc *ProductController) CreateProduct(c echo.Context) error {
	var createProductRequest request.CreateProductRequest

	if err := c.Bind(&createProductRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error parsing JSON",
		})
	}

	productCommand, err := createProductRequest.ToCreateProductCommand()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid Product",
		})
	}

	result, err := pc.service.CreateProduct(productCommand)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error when creating product",
		})
	}

	return c.JSON(http.StatusCreated, result.Result)
}

func (pc *ProductController) GetAllProducts(c echo.Context) error {
	products, err := pc.service.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error getting products",
		})
	}

	return c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProductByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product",
		})
	}

	product, err := pc.service.FindProductByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error getting the product by id",
		})
	}

	return c.JSON(http.StatusOK, product)
}
