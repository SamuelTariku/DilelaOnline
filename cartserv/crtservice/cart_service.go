package crtservice

import(
	"../../entity"
	"../../cartserv"
)

type CartService struct {
	cartRepo cartserv.CartRepository
}

func NewCartService(CartRepo cartserv.CartRepository) *CartService{
	return &CartService{cartRepo:CartRepo}
}

func (c *CartService) Carts() ([]entity.Cart, error){
	car, err := c.cartRepo.Carts()
	if err != nil {
		return nil, err
	}
	return car, nil
}



func (c *CartService) Cart(id int) (entity.Cart, error) {
	cart, err := c.cartRepo.Cart(id)
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (c *CartService) UserCart(userid int) ([]entity.Cart, error){
	cart, err := c.cartRepo.UserCart(userid)
	if err != nil{
		return nil, err
	}
	return cart,nil
}

func (c *CartService) UpdateC(car entity.Cart) error {
	err := c.cartRepo.UpdateC(car)
	if err != nil {
		return err

	}
	return nil
}

func (c *CartService) StoreC(car entity.Cart) error {
	err := c.cartRepo.StoreC(car)
	if err != nil {
		return err
	}
	return nil
}

func (c *CartService) DeleteC(id int) error {
	err := c.cartRepo.DeleteC(id)
	if err != nil {
		return err

	}
	return nil
}