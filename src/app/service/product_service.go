package service

import (
	"GolangwithFrame/src/domain/model"
	"GolangwithFrame/src/infrastructure/repository"
)

type ProductService interface {
	CreateProduct(model.Product) model.Product
	FindAllProducts() []model.Product
	UpdateProduct(product model.Product) error
	DeleteProduct(product model.Product) error
	GetProduct(id int) (model.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func New(repo repository.ProductRepository) ProductService {
	return &productService{
		productRepository: repo,
	}
}

//type Service struct {
//	productRepository repository.ProductRepository
//}
//func New(repo repository.Repository) Service {
//	return &Service{
//		Repository: repo,
//	}
//}

func (service *productService) CreateProduct(product model.Product) model.Product {
	service.productRepository.CreateProduct(product)
	return product
}

func (service *productService) FindAllProducts() []model.Product {
	return service.productRepository.FindAllProducts()
}

func (service *productService) UpdateProduct(product model.Product) error {
	return service.productRepository.UpdateProduct(product)
}

func (service *productService) DeleteProduct(product model.Product) error {
	return service.productRepository.DeleteProduct(product)

}

func (service *productService) GetProduct(id int) (model.Product, error) {
	return service.productRepository.GetProduct(id)
}
