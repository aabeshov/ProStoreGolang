package main

import (
	"GolangwithFrame/app/controller"
	middleware2 "GolangwithFrame/app/middleware"
	"GolangwithFrame/src/app/service"
	"GolangwithFrame/src/infrastructure/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

var (
	productRepository repository.ProductRepository = repository.NewProductRepository()
	productService    service.ProductService       = service.New(productRepository)
	productController controller.ProductController = controller.New(productService)
)

func main() {
	defer productRepository.CloseDB()
	server := gin.Default()
	server.Use(
		gin.Recovery(),
		middleware2.Logger(),
		middleware2.BasicAuth(),
		//gindump.Dump()
	)

	//database.Init()

	server.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, productController.FindAll())
	})
	server.POST("/products", func(ctx *gin.Context) {
		err := productController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Product was created"})

		}
	})
	server.PUT("/products/:id", func(ctx *gin.Context) {
		err := productController.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Product was updated"})

		}
	})

	server.DELETE("/products/:id", func(ctx *gin.Context) {
		err := productController.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Product was deleted!"})

		}

	})
	server.GET("/products/:id", func(ctx *gin.Context) {
		prod, err := productController.FindById(ctx)
		if err != nil {
			panic("error")
		} else {
			ctx.JSON(200, gin.H{"message": prod})
		}

	})
	server.Run(":8080")

}
