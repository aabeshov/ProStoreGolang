package service

import (
	"GolangwithFrame/src/domain/model"
	"GolangwithFrame/src/infrastructure/repository"
)

type ProductService interface {
	Save(model.Product) model.Product
	FindAll() []model.Product
	Update(product model.Product)
	Delete(product model.Product)
	Get(id int) (model.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func New(repo repository.ProductRepository) ProductService {
	return &productService{
		productRepository: repo,
	}
}

func (service *productService) Save(product model.Product) model.Product {
	service.productRepository.Save(product)
	return product
}

func (service *productService) FindAll() []model.Product {
	return service.productRepository.FindAll()
}

func (service *productService) Update(product model.Product) {
	service.productRepository.Update(product)
}

func (service *productService) Delete(product model.Product) {
	service.productRepository.Delete(product)
}

func (service *productService) Get(id int) (model.Product, error) {
	return service.productRepository.Get(id)
}
