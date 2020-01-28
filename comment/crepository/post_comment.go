package crepository

import (
	"database/sql"
	"errors"
	"log"

	"../../entity"
)

//
type PostCommRepository struct {
	conn *sql.DB
}

//
func NewCommPostRepo(Conn *sql.DB) *PostCommRepository {
	return &PostCommRepository{conn: Conn}
}

//Return all comments
func (commRepo *PostCommRepository) Comments() ([]entity.Comment, error) {

	rows, err := commRepo.conn.Query("SELECT * FROM comments;")
	if err != nil {
		return nil, errors.New("could not query the database")
	}
	defer rows.Close()

	comms := []entity.Comment{}

	for rows.Next() {
		comm := entity.Comment{}
		err = rows.Scan(&comm.ID, &comm.Name, &comm.Message, &comm.Email, &comm.PlacedAt, &comm.UserID, &comm.ProductID, &comm.Rating)
		if err != nil {
			return nil, err
		}
		comms = append(comms, comm)
	}

	return comms, nil
}

//
func (commRepo *PostCommRepository) Comment(id int) (entity.Comment, error) {
	row := commRepo.conn.QueryRow("SELECT * FROM comments WHERE id = $1", id)

	u := entity.Comment{}

	err := row.Scan(&u.ID, &u.Name, &u.Message, &u.Email, &u.PlacedAt, &u.UserID, &u.ProductID, &u.Rating)

	log.Println(u.Name)
	log.Println(u.Message)
	log.Println(u.Email)
	log.Println(u.PlacedAt)
	log.Println(u.UserID)
	log.Println(u.ProductID)
	log.Println(u.Rating)

	if err != nil {
		return u, err
	}

	return u, nil

}

//Find comments for product
func (commRepo *PostCommRepository) ProductComment(productid int) ([]entity.Comment, error) {
	rows, err := commRepo.conn.Query("SELECT * FROM comments WHERE productid = $1", productid)
	if err != nil {
		return nil, errors.New("could not query the database")
	}
	defer rows.Close()

	comms := []entity.Comment{}

	for rows.Next() {
		comm := entity.Comment{}
		//HOTFIX AGAIN!!

		err = rows.Scan(&comm.ID, &comm.Name, &comm.Email, &comm.PlacedAt, &comm.ProductID, &comm.Rating, &comm.Message,&comm.UserID)
		if err != nil {

		}
		comms = append(comms, comm)
	}

	if err != nil {
		panic(err)
	}

	return comms, nil
}

//
func (commRepo *PostCommRepository) UpdateComment(u entity.Comment) error {

	_, err := commRepo.conn.Exec("UPDATE comments SET name=$1,message=$2,email=$3,commTime=$4,userid=$5, productid=$6, rating=$7,  WHERE id=$5",
		u.Name, u.Message, u.Email, u.PlacedAt, u.UserID, u.ProductID, u.Rating, u.ID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

//
func (commRepo *PostCommRepository) DeleteComment(id int) error {

	_, err := commRepo.conn.Exec("DELETE FROM comments WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

//
func (commRepo *PostCommRepository) StoreComment(u entity.Comment) error {
	if(u.UserID == 0){
		_, err := commRepo.conn.Exec("INSERT INTO comments (name, message, email, productid, rating, commtime)"+
			" values($1, $2, $3, $4, $5,current_timestamp)", u.Name, u.Message, u.Email, u.ProductID, u.Rating)
		if err != nil {
			//panic(err)
			return errors.New("Insertion has failed")
		}
	} else {
		_, err := commRepo.conn.Exec("INSERT INTO comments (name, message, email, userid, productid, rating, commtime)"+
			" values($1, $2, $3, $4, $5, $6,current_timestamp)", u.Name, u.Message, u.Email, u.UserID, u.ProductID, u.Rating)
		if err != nil {
			//panic(err)
			return errors.New("Insertion has failed")
		}
	}


	return nil
}
