package crtrepository

import (
	"../../entity"
	"database/sql"
	"errors"
)

type PostCartRepo struct {
	conn *sql.DB
}

func NewPostCartRepo(Conn *sql.DB) *PostCartRepo{
	return &PostCartRepo{conn: Conn}
}

func(c *PostCartRepo) Carts() ([]entity.Cart, error){
	rows, err := c.conn.Query("SELECT * from cart")
	if err != nil{
		return nil, errors.New("could not query")
	}
	defer rows.Close()
	car := []entity.Cart{}
	for rows.Next(){
		cart := entity.Cart{}
		err = rows.Scan(&cart.ID,&cart.ProductID,&cart.UserID,&cart.Price,&cart.AddedTime,&cart.ProductName)
		if err != nil{
			return nil, err
		}
		car = append(car, cart)
	}
	return car, nil
}
func (c *PostCartRepo) UserCart(userid int) ([]entity.Cart, error){
	rows, err := c.conn.Query("SELECT * from cart where userid = $1", userid)
	if err != nil{
		return nil, errors.New("could not query")
	}
	defer rows.Close()
	car := []entity.Cart{}
	for rows.Next(){
		cart := entity.Cart{}
		err = rows.Scan(&cart.ID,&cart.ProductID,&cart.UserID,&cart.Price,&cart.AddedTime, &cart.ProductName)
		if err !=nil {
			return nil, err
		}
		car = append(car,cart)
	}
	return car, nil
}
func (c *PostCartRepo) Cart(id int) (entity.Cart, error) {
	rows := c.conn.QueryRow("SELECT * from cart WHERE id = $1", id)

	car := entity.Cart{}

	err := rows.Scan(&car.ID, &car.ProductID, &car.UserID, &car.Price, &car.AddedTime, &car.ProductName)

	if err != nil {
		return car, err
	}
	return car, nil

}

func (c *PostCartRepo) UpdateC(car entity.Cart) error {
	_, err := c.conn.Exec("UPDATE cart SET productid = $1, userid=$2, price=$3, productname=$4 WHERE id = $5", car.ProductID, car.UserID, car.Price,car.ProductName, car.ID)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

func (c *PostCartRepo) StoreC(car entity.Cart) error {
	_, err := c.conn.Exec("INSERT INTO cart (productid,userid,price,productname,addedtime)"+"values($1,$2,$3,$4,current_timestamp)", car.ProductID, car.UserID, car.Price, car.ProductName)
	if err != nil {
		panic(err)
		return errors.New("failed to store")
	}
	return nil
}

func (c *PostCartRepo) DeleteC(id int) error {
	_, err := c.conn.Exec("DELETE FROM cart WHERE id=$1", id)
	if err != nil {
		return errors.New("failed to delete")
	}
	return nil
}