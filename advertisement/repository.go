package advertisement

import (
	"../entity"
)

type AdvertRepo interface {
	Adverts() ([]entity.Advertisement, error)
	Advert(id int) (entity.Advertisement, error)
	UpdateA(a entity.Advertisement) error
	DeleteA(id int) error
	StoreA(a entity.Advertisement) error
}
