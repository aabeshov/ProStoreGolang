package main

import (
	"GolangwithFrame/app/controller"
	middleware2 "GolangwithFrame/app/middleware"
	"GolangwithFrame/src/app/service"
	"GolangwithFrame/src/infrastructure/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	productRepository repository.ProductRepository = repository.NewRepository()
	productService    service.ProductService       = service.New(productRepository)
	productController controller.ProductController = controller.New(productService)

	//Repository repository.ProductRepository = repository.New()
	//Service    service.SER                  = service.New(Repository)
	//Controller controller.ProductController = controller.New(Service)
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

	//cart := server.Group("/cart")
	//{
	//	cart.GET("/:id", Controller.SelectFromCartByID)
	//	cart.DELETE("/:id", Controller.DeleteFromCartByID)
	//	cart.POST("", Controller.AddToCart)
	//	cart.PUT("", Controller.UpdateCartControl)
	//}

	products := server.Group("/products")
	{
		products.GET("", productController.FindAllProducts)
		products.POST("", productController.CreateProduct)
		products.GET("/:id", productController.GetProduct)
		products.PUT("/:id", productController.UpdateProduct)
		products.DELETE("/:id", productController.DeleteProduct)
	}

	//server.GET("/products", func(ctx *gin.Context) {
	//	ctx.JSON(200, productController.FindAll)
	//})
	//server.POST("/products", func(ctx *gin.Context) {
	//	err := productController.Save(ctx)
	//	if err != nil {
	//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	} else {
	//		ctx.JSON(http.StatusOK, gin.H{"message": "Product was created"})
	//
	//	}
	//})
	//server.PUT("/products/:id", func(ctx *gin.Context) {
	//	err := productController.Update(ctx)
	//	//fmt.Println(err)
	//	if err != nil {
	//		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		if err.Error() == "record not found" {
	//			ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is nothing to update"})
	//		} else {
	//			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Used invalid ID"})
	//		}
	//	} else {
	//		ctx.JSON(http.StatusOK, gin.H{"message": "Product was updated"})
	//
	//	}
	//})
	//
	//server.DELETE("/products/:id", func(ctx *gin.Context) {
	//	err := productController.Delete(ctx)
	//	//fmt.Println(err)
	//	if err != nil {
	//		if err.Error() == "record not found" {
	//			ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is nothing to delete"})
	//
	//		} else {
	//			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Used invalid ID"})
	//		}
	//		//fmt.Println(err.Error()[0:17])
	//	} else {
	//		ctx.JSON(http.StatusOK, gin.H{"message": "Product was deleted!"})
	//
	//	}
	//})
	//server.GET("/products/:id", func(ctx *gin.Context) {
	//	prod, err := productController.FindById(ctx)
	//	//fmt.Println(err, err.Error())
	//	if err != nil {
	//		if err.Error() == "record not found" {
	//			ctx.JSON(404, gin.H{"message": "There is no object with this ID"})
	//		} else {
	//			ctx.JSON(404, gin.H{"message": "Used invalid ID"})
	//		}
	//	} else {
	//		ctx.JSON(200, gin.H{"message": prod})
	//	}
	//})
	server.Run(":8080")

}
