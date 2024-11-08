package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigospimentacwb/go-gin-rest/dto"
	"github.com/rodrigospimentacwb/go-gin-rest/model"
	"github.com/rodrigospimentacwb/go-gin-rest/usecase"
	"net/http"
	"strconv"
)

type ProductController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(useCase usecase.ProductUseCase) ProductController {
	return ProductController{
		productUseCase: useCase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *ProductController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := dto.Response{Message: "product id required"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := dto.Response{Message: "product ID must be a number."}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := dto.Response{Message: "product not found"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
