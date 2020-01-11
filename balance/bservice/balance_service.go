package bservice

import (
	"../../balance"
	"../../entity"
)

type BalanceService struct {
	balanceRepo balance.BalanceRepo
}

func NewBalanceService(BalRepo balance.BalanceRepo) *BalanceService {
	return &BalanceService{balanceRepo: BalRepo}
}

func (b *BalanceService) Balance(id int) (entity.Balance, error) {
	bal, err := b.balanceRepo.Balance(id)
	if err != nil {
		return bal, err
	}
	return bal, nil
}

func (b *BalanceService) Updateb(bal entity.Balance) error {
	err := b.balanceRepo.Updateb(bal)
	if err != nil {
		return err

	}
	return nil
}

func (b *BalanceService) Storeb(id int, bal entity.Balance) error {
	err := b.balanceRepo.Storeb(id, bal)
	if err != nil {
		return err
	}
	return nil
}
func (b *BalanceService) StoreId(id uint) error {
	err := b.balanceRepo.StoreId(id)
	if err != nil {
		return err

	}
	return nil
}

func (b *BalanceService) Deleteb(id int) error {
	err := b.balanceRepo.Deleteb(id)
	if err != nil {
		return err

	}
	return nil
}
