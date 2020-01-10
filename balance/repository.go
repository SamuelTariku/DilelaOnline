package balance

import (
	"../entity"
)

type BalanceRepo interface {
	Balance(id int) (entity.Balance, error)
	Updateb(p entity.Balance) error
	Deleteb(id int) error
	Storeb(b entity.Balance) error
}
