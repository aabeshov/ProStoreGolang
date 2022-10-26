package controller

import (
	"GolangwithFrame/src/app/service"
	"GolangwithFrame/src/domain/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ProductController interface {
	FindAll() []model.Product
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	FindById(ctx *gin.Context) (model.Product, error)
}

type controller struct {
	service service.ProductService
}

func New(service service.ProductService) ProductController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []model.Product {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var product model.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		return err
	}

	c.service.Save(product)
	return nil

}

func (c *controller) Update(ctx *gin.Context) error {
	var product model.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	product.Id = id
	c.service.Update(product)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var product model.Product
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	product.Id = id
	c.service.Delete(product)
	return nil
}

func (c *controller) FindById(ctx *gin.Context) (model.Product, error) {
	ProductId := ctx.Param("id")
	ProductIdInt, err := strconv.Atoi(ProductId)
	if err != nil {
		panic("Error")
	}
	var product model.Product
	ctx.ShouldBindJSON(&product)

	return c.service.Get(ProductIdInt)

}
