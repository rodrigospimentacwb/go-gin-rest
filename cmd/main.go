package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigospimentacwb/go-gin-rest/controller"
	"github.com/rodrigospimentacwb/go-gin-rest/db"
	"github.com/rodrigospimentacwb/go-gin-rest/repository"
	"github.com/rodrigospimentacwb/go-gin-rest/usecase"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":8080")
}
