package balance

import (
	"../entity"
)

type BalanceBalance interface {
	Balance(id int) (entity.Balance, error)
	Updateb(p entity.Balance)
	Deleteb(id int) error
	Storeb(b entity.Balance) error
}
