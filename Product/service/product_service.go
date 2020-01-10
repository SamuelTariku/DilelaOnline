package service

import (
	"../../entity"
	"../../product"
)

type ProductService struct {
	productRepo product.ProductRepo
}

func NewProductService(ProRepo product.ProductRepo) *ProductService {
	return &ProductService{productRepo: ProRepo}
}

func (p *ProductService) Products() ([]entity.Product, error) {
	prod, err := p.productRepo.Products()
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func (p *ProductService) Product(id int) (entity.Product, error) {
	prod, err := p.productRepo.Product(id)
	if err != nil {
		return prod, err
	}
	return prod, nil
}

func (p *ProductService) UpdateP(prod entity.Product) error {
	err := p.productRepo.UpdateP(prod)
	if err != nil {
		return err

	}
	return nil
}

func (p *ProductService) StoreP(prod entity.Product) error {
	err := p.productRepo.StoreP(prod)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) DeleteP(id int) error {
	err := p.productRepo.DeleteP(id)
	if err != nil {
		return err

	}
	return nil
}

func (p *ProductService) SearchProduct(prod string) ([]entity.Product, error) {
	products, err := p.productRepo.SearchProduct(prod)
	if err != nil {
		return nil, err
	}
	return products, nil
}
