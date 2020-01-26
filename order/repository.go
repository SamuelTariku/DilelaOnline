package order

import (
	"../entity"
)

type OrderRepo interface {
	Orders() ([]entity.Order, error)
	Order(id int) (entity.Order, error)
	UpdateO(p entity.Order) error
	DeleteO(id int) error
	StoreO(p entity.Order) error
}