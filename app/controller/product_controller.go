package controller

import (
	"GolangwithFrame/src/app/service"
	"GolangwithFrame/src/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController interface {
	FindAllProducts(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	GetProduct(ctx *gin.Context)
}

type controllerProduct struct {
	service service.ProductService
}

func New(service service.ProductService) ProductController {
	return &controllerProduct{
		service: service,
	}
}

func (c *controllerProduct) FindAllProducts(ctx *gin.Context) {
	ctx.JSON(200, c.service.FindAllProducts())
}

func (c *controllerProduct) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.service.CreateProduct(product)
	ctx.JSON(http.StatusOK, gin.H{"message": "Product was created"})

}

func (c *controllerProduct) UpdateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Used invalid ID"})
	}
	product.Id = id
	err = c.service.UpdateProduct(product)
	//fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is nothing to update"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product was updated"})

}

func (c *controllerProduct) DeleteProduct(ctx *gin.Context) {
	var product model.Product
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Used invalid ID"})
	}
	product.Id = id
	err = c.service.DeleteProduct(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is nothing to delete"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product was deleted!"})

}

func (c *controllerProduct) GetProduct(ctx *gin.Context) {
	var product model.Product
	ProductId := ctx.Param("id")
	ProductIdInt, err := strconv.Atoi(ProductId)
	if err != nil {
		ctx.JSON(404, gin.H{"message": "Used invalid ID"})
	}
	ctx.ShouldBindJSON(&product)
	prod, err := c.service.GetProduct(ProductIdInt)
	if err != nil {
		ctx.JSON(404, gin.H{"message": "There is no object with this ID"})

	}
	ctx.JSON(200, gin.H{"message": prod})
	//fmt.Println(err, err.Error())

}
