package main

import (
	"GolangwithFrame/app/controller"
	middleware2 "GolangwithFrame/app/middleware"
	"GolangwithFrame/src/app/service"
	"GolangwithFrame/src/infrastructure/localvariables"
	"GolangwithFrame/src/infrastructure/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	Repository = repository.NewRepository()
	Service    = service.New(Repository)
	controller = Controller.New(*Service)

	//Repository repository.ProductRepository = repository.New()
	//Service    service.SER                  = service.New(Repository)
	//controller controller.ProductController = controller.New(Service)
)

func init() {
	localvariables.LoadEnvVariables()
}

func main() {
	defer Repository.CloseDB()
	//server := gin.New()
	server := gin.Default()
	server.Use(
		//gin.Recovery(),
		//middleware2.Logger(),
		//middleware2.BasicAuth(),
	)

	products := server.Group("/products", middleware2.RequireAuth)
	//products.Use(middleware2.RequireAuth)
	{
		products.GET("", controller.FindAllProducts)
		products.POST("", controller.CreateProduct)
		products.GET("/:id", controller.GetProduct)
		products.PUT("/:id", controller.UpdateProduct)
		products.DELETE("/:id", controller.DeleteProduct)
	}

	category := server.Group("/category", middleware2.RequireAuth)
	//category.Use(middleware2.RequireAuth)
	{

		category.GET("", controller.FindAllCategory)
		category.GET("/:id/products", controller.FindProductsByCategory)
		category.POST("", controller.CreateCategory)
		category.GET("/:id", controller.GetCategory)
		category.PUT("/:id", controller.UpdateCategory)
		category.DELETE("/:id", controller.DeleteCategory)
	}

	users := server.Group("/account")
	{
		users.POST("/signup", controller.SignUp)
		users.POST("/login", controller.Login)
		users.GET("/validate", middleware2.RequireAuth, controller.Validate)

	}

	server.Run(":8080")

	//cart := server.Group("/cart")
	//{
	//	cart.GET("/:id", controller.SelectFromCartByID)
	//	cart.DELETE("/:id", controller.DeleteFromCartByID)
	//	cart.POST("", controller.AddToCart)
	//	cart.PUT("", controller.UpdateCartControl)
	//}

	//server.GET("/products", func(ctx *gin.Context) {
	//	ctx.JSON(200, controller.FindAll)
	//})
	//server.POST("/products", func(ctx *gin.Context) {
	//	err := controller.Save(ctx)
	//	if err != nil {
	//		ctx.JSON(http.StatusBadRequest, gin.H{"error": erгr.Error()})
	//	} else {
	//		ctx.JSON(http.StatusOK, gin.H{"message": "Product was created"})
	//
	//	}
	//})
	//server.PUT("/products/:id", func(ctx *gin.Context) {
	//	err := controller.Update(ctx)
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
	//	err := controller.Delete(ctx)
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
	//	prod, err := controller.FindById(ctx)
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
}
