package repository

import (
	"../../entity"
	"database/sql"
	"errors"
	//"log"
)

type PostBalanceRepo struct {
	conn *sql.DB
}

func NewBalanceRepo(Conn *sql.DB) *PostBalanceRepo {
	return &PostBalanceRepo{conn: Conn}
}

func (b *PostBalanceRepo) Balance(id int) (entity.Balance, error) {

	row := b.conn.QueryRow("SELECT * FROM balance WHERE id=$1", id)
	balance := entity.Balance{}
	err := row.Scan(&balance.ID, &balance.YourBalance)

	if err != nil {
		return balance, err
	}
	return balance, nil
}

func (b *PostBalanceRepo) Deleteb(id int) error {
	_, err := b.conn.Exec("DELETE FROM balance WHERE id=$1", id)
	if err != nil {
		return errors.New("failed to delete")
	}
	return nil
}

func (b *PostBalanceRepo) Updateb(bal entity.Balance) error {
	_, err := b.conn.Exec("UPDATE balance SET id = $1, balance = $2", bal.ID, bal.YourBalance)
	if err != nil {
		return errors.New("failed to update")
	}
	return nil
}
func (b *PostBalanceRepo) Storeb(bal entity.Balance) error {
	_, err := b.conn.Exec("INSERT INTO balance (id,balance)"+"values($1,$2)", bal.ID, bal.YourBalance)
	if err != nil {
		panic(err)
		return errors.New("failed to store balance")
	}
	return nil
}
