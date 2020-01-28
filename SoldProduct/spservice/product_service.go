package spservice

import (
	"../../entity"
	"../../SoldProduct"
)

type ProductService struct {
	soldProductRepo SoldProduct.ProductRepo
}

func NewProductService(ProRepo SoldProduct.ProductRepo) *ProductService {
	return &ProductService{soldProductRepo: ProRepo}
}

func (p *ProductService) Products() ([]entity.Product, error) {
	prod, err := p.soldProductRepo.Products()
	if err != nil {
		return nil, err
	}
	return prod, err
}

func (p *ProductService) SoldP(id int) ([]entity.Product, error) {
	prod, err := p.soldProductRepo.SoldP(id)
	if err != nil {
		return nil, err
	}
	return prod, err
}
func (p *ProductService) Product(id int) (entity.Product, error) {
	prod, err := p.soldProductRepo.Product(id)
	if err != nil {
		return prod, err
	}
	return prod, err
}

func (p *ProductService) UpdateP(prod entity.Product) error {
	err := p.soldProductRepo.UpdateP(prod)
	if err != nil {
		return err

	}
	return err
}

func (p *ProductService) StoreP(prod entity.Product) error {
	err := p.soldProductRepo.StoreP(prod)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) DeleteP(id int) error {
	err := p.soldProductRepo.DeleteP(id)
	if err != nil {
		return err

	}
	return err
}

func (p *ProductService) SearchProduct(prod string) ([]entity.Product, error) {
	products, err := p.soldProductRepo.SearchProduct(prod)
	if err != nil {
		return nil, err
	}
	return products, err
}
