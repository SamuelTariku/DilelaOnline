package product

import (
	"../entity"
)

type SProductRepo interface {
	Products() ([]entity.Product, error)
	Product(id int) (entity.Product, error)
	UpdateP(p entity.Product) error
	DeleteP(id int) error
	StoreP(p entity.Product) error
	SearchProduct(prod string) ([]entity.Product, error)
}
