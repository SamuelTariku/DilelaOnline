package oservice

import (
	"../../entity"
	"../../order"
)

type OrderService struct {
	orderRepo order.OrderRepo
}

func NewOrderService(OrdRepo order.OrderRepo) *OrderService {
	return &OrderService{orderRepo: OrdRepo}
}

func (o *OrderService) Orders() ([]entity.Order, error) {
	ord, err := o.orderRepo.Orders()
	if err != nil {
		return nil, err
	}
	return ord, nil
}

func (o *OrderService) Order(id int) (entity.Order, error) {
	ord, err := o.orderRepo.Order(id)
	if err != nil {
		return ord, err
	}
	return ord, nil
}

func (o *OrderService) UpdateO(ord entity.Order) error {
	err := o.orderRepo.UpdateO(ord)
	if err != nil {
		return err
	}
	return err
}

func (o *OrderService) StoreO(ord entity.Order) error {
	err := o.orderRepo.StoreO(ord)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderService) DeleteO(id int) error {
	err := o.orderRepo.DeleteO(id)
	if err != nil {
		return err
	}
	return nil
}
