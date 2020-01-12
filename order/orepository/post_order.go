package orepository

import (
	"../../entity"
	"database/sql"
	"errors"
)

type PostOrderRepo struct {
	conn *sql.DB
}

func NewOrderRepo(Conn *sql.DB) *PostOrderRepo {
	return &PostOrderRepo{conn: Conn}
}

func (o *PostOrderRepo) Orders() ([]entity.Order, error) {
	rows, err := o.conn.Query("SELECT * from order;")
	if err != nil {
		return nil, errors.New("could not query")
	}
	defer rows.Close()
	ord := []entity.Order{}

	for rows.Next() {
		order := entity.Order{}
		err = rows.Scan(&order.ID, &order.PlacedAt, &order.UserID, &order.ItemID)
		if err != nil {
			return nil, err
		}
		ord = append(ord, order)
	}
	return ord, nil
}

func (o *PostOrderRepo) Order(id int) (entity.Order, error) {
	rows := o.conn.QueryRow("SELECT * FROM order WHERE id = $1", id)
	ord := entity.Order{}
	err := rows.Scan(&ord.ID, &ord.PlacedAt, &ord.UserID, &ord.ItemID)
	if err != nil {
		return ord, err
	}
	return ord, nil
}

func (o *PostOrderRepo) UpdateO(ord entity.Order) error {
	_, err := o.conn.Exec("UPDATE order SET userid = $1, itemid = $2", ord.UserID, ord.ItemID)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

func (o *PostOrderRepo) StoreO(ord entity.Order) error {
	_, err := o.conn.Exec("INSERT INTO order (userid,itemid)"+"values($1,$2)", ord.UserID, ord.ItemID)
	if err != nil {
		panic(err)
		return errors.New("failed to store")
	}
	return nil
}

func (o *PostOrderRepo) DeleteO(id int) error {
	_, err := o.conn.Exec("DELETE FROM order WHERE id = $1", id)
	if err != nil {
		return errors.New("failed to delete")
	}
	return nil
}
