package product

import (
	"../entity"
)

type ProductService interface {
	Products() ([]entity.Product, error)
	Product(id int) (entity.Product, error)
	Bytype(typ string) (entity.Product, error)
	UpdateP(p entity.Product) error
	DeleteP(id int) error
	StoreP(p entity.Product) error
	SearchProduct(prod string) ([]entity.Product, error)
}
