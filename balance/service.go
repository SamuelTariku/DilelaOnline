package balance

import "../entity"

type BalanceService interface {
	Balance(id int) (entity.Balance, error)
	Updateb(p entity.Balance) error
	Deleteb(id int) error
	Storeb(id int, b entity.Balance) error
	StoreId(id uint) error
}
