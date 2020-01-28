package arepository

import (
	"../../entity"
	"database/sql"
	"errors"
)

type PostAdvertRepo struct {
	conn *sql.DB
}

func NewPostAdvertRepo(Conn *sql.DB) *PostAdvertRepo {
	return &PostAdvertRepo{conn: Conn}
}

func (a *PostAdvertRepo) Adverts() ([]entity.Advertisement, error) {
	rows, err := a.conn.Query("SELECT * from advert;")
	if err != nil {
		return nil, errors.New("could not query")
	}
	defer rows.Close()
	ad := []entity.Advertisement{}
	for rows.Next() {
		advertisement := entity.Advertisement{}
		err = rows.Scan(&advertisement.ID, &advertisement.ProductID, &advertisement.Ptype)
		if err != nil {
			return nil, err
		}
		ad = append(ad, advertisement)
	}
	return ad, nil
}

func (a *PostAdvertRepo) Advert(id int) (entity.Advertisement, error) {
	rows := a.conn.QueryRow("SELECT * from advert WHERE id = $1", id)

	ad := entity.Advertisement{}

	err := rows.Scan(&ad.ID, &ad.ProductID, &ad.Ptype)

	if err != nil {
		return ad, err
	}
	return ad, nil

}

func (a *PostAdvertRepo) UpdateA(ad entity.Advertisement) error {
	_, err := a.conn.Exec("UPDATE advert productid = $1, ptype=$2 WHERE id = $6", ad.ProductID, ad.Ptype, ad.ID)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

func (a *PostAdvertRepo) StoreA(ad entity.Advertisement) error {
	_, err := a.conn.Exec("INSERT INTO advert (productid,ptype)"+"values($1,$2)", ad.ProductID, ad.Ptype)
	if err != nil {
		panic(err)
		return errors.New("failed to store")
	}
	return nil
}

func (a *PostAdvertRepo) DeleteA(id int) error {
	_, err := a.conn.Exec("DELETE FROM advert WHERE id=$1", id)
	if err != nil {
		return errors.New("failed to delete")
	}
	return nil
}
