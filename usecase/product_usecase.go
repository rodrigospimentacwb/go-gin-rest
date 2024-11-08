package usecase

import (
	"github.com/rodrigospimentacwb/go-gin-rest/model"
	"github.com/rodrigospimentacwb/go-gin-rest/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = productId
	return product, nil
}

func (pu *ProductUseCase) GetProductById(productId int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}
