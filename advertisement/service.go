package advertisement

import (
	"../entity"
)

type AdvertService interface {
	Adverts() ([]entity.Advertisement, error)
	Advert(id int) (entity.Advertisement, error)
	UpdateA(a entity.Advertisement) error
	DeleteA(id int) error
	StoreA(a entity.Advertisement) error
}
