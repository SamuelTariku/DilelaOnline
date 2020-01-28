package prepository

import (
	"../../entity"
	"database/sql"
	"errors"
)

type PostProductRepo struct {
	conn *sql.DB
}

func NewPostProductRepo(Conn *sql.DB) *PostProductRepo {
	return &PostProductRepo{conn: Conn}
}

func (p *PostProductRepo) Products() ([]entity.Product, error) {
	rows, err := p.conn.Query("SELECT * FROM product;")
	if err != nil {
		return nil, errors.New("Could not query")
	}
	defer rows.Close()
	prod := []entity.Product{}

	for rows.Next() {
		product := entity.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Ptype, &product.Price, &product.Description, &product.CreatedAt, &product.Image, &product.UserID)
		if err != nil {
			return nil, err
		}
		prod = append(prod, product)
	}

	return prod, nil
}

func (p *PostProductRepo) Product(id int) (entity.Product, error) {
	rows := p.conn.QueryRow("SELECT * from product WHERE id = $1", id)

	prod := entity.Product{}

	err := rows.Scan(&prod.ID, &prod.Name, &prod.Ptype, &prod.Price, &prod.Description, &prod.CreatedAt, &prod.Image, &prod.UserID)


	if err != nil {
		return prod, err
	}
	return prod, nil

}

func (p *PostProductRepo) Bytype(typ string) (entity.Product, error){
	rows := p.conn.QueryRow("Select * from product Where ptype = $1", typ)

	prod := entity.Product{}
	err := rows.Scan(&prod.ID, &prod.Name, &prod.Ptype, &prod.Price, &prod.Description, &prod.CreatedAt, &prod.Image, &prod.UserID)


	if err != nil {
		return prod, err
	}
	return prod, nil
}
func (p *PostProductRepo) UpdateP(pro entity.Product) error {
	_, err := p.conn.Exec("UPDATE product SET name = $1, ptype=$2, price=$3, description=$4,Image=$5,userid=$6 WHERE id = $6", pro.Name, pro.Ptype, pro.Price, pro.Description, pro.Image, pro.ID, pro.UserID)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

func (p *PostProductRepo) StoreP(pro entity.Product) error {
	_, err := p.conn.Exec("INSERT INTO product (name,ptype,price,description,Image,userid,createdat)"+"values($1,$2,$3,$4,$5,$6,current_timestamp)", pro.Name, pro.Ptype, pro.Price, pro.Description, pro.Image, pro.UserID)
	if err != nil {
		panic(err)
		return errors.New("failed to store")
	}
	return nil
}

func (p *PostProductRepo) DeleteP(id int) error {
	_, err := p.conn.Exec("DELETE FROM product WHERE id=$1", id)
	if err != nil {
		return errors.New("failed to delete")
	}
	return nil
}

func (p *PostProductRepo) SearchProduct(prod string) ([]entity.Product, error) {
	rows, err := p.conn.Query("SELECT * FROM product WHERE name LIKE $1", "%"+prod+"%")

	if err != nil {
		errors.New("could not query")
	}
	defer rows.Close()

	pr := []entity.Product{}
	for rows.Next() {
		product := entity.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Ptype, &product.Price, &product.Description, &product.CreatedAt, &product.Image, &product.UserID)
		if err != nil {
			return nil, err

		}
		pr = append(pr, product)
	}
	return pr, nil
}
