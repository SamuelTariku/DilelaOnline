package cartserv

import(
	"../entity"
)

type CartRepository interface {
	Carts() ([]entity.Cart,error)
	UserCart(userid int) ([]entity.Cart,error)
	Cart(id int) (entity.Cart, error)
	UpdateC(c entity.Cart) error
	DeleteC(id int) error
	StoreC(c entity.Cart) error
}
