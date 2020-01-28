package aservice

import (
	"../../advertisement"
	"../../entity"
)

type AdvertService struct {
	advertRepo advertisement.AdvertRepo
}

func NewAdvertService(adRepo advertisement.AdvertRepo) *AdvertService {
	return &AdvertService{advertRepo: adRepo}
}

func (a *AdvertService) Adverts() ([]entity.Advertisement, error) {
	ad, err := a.advertRepo.Adverts()
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (a *AdvertService) Advert(id int) (entity.Advertisement, error) {
	ad, err := a.advertRepo.Advert(id)
	if err != nil {
		return ad, err
	}
	return ad, nil
}

func (a *AdvertService) UpdateA(ad entity.Advertisement) error {
	err := a.advertRepo.UpdateA(ad)
	if err != nil {
		return err

	}
	return nil
}

func (a *AdvertService) StoreA(ad entity.Advertisement) error {
	err := a.advertRepo.StoreA(ad)
	if err != nil {
		return err
	}
	return nil
}

func (a *AdvertService) DeleteA(id int) error {
	err := a.advertRepo.DeleteA(id)
	if err != nil {
		return err

	}
	return nil
}
