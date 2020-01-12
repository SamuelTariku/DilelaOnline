package repository

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

//
func (commRepo *PostCommRepository) Comments() ([]entity.Comment, error) {

	rows, err := commRepo.conn.Query("SELECT * FROM comments;")
	if err != nil {
		return nil, errors.New("could not query the database")
	}
	defer rows.Close()
	comms := []entity.Comment{}

	for rows.Next() {
		comm := entity.Comment{}
		err = rows.Scan(&comm.ID, &comm.UserID, &comm.Name, &comm.Message, &comm.Email, &comm.Rating, &comm.PlacedAt)
		if err != nil {
			return nil, err
		}
		comms = append(comms, comm)
	}

	return comms, nil
}

//
func (commRepo *PostCommRepository) CommentWithID(id int) (entity.Comment, error) {
	row := commRepo.conn.QueryRow("SELECT * FROM comments WHERE id = $1", id)

	u := entity.Comment{}

	err := row.Scan(&u.ID, &u.Name, &u.Message, &u.Email, &u.PlacedAt)

	log.Println(u.UserID)
	log.Println(u.Name)
	log.Println(u.Message)
	log.Println(u.Email)
	log.Println(u.Rating)
	log.Println(u.PlacedAt)

	if err != nil {
		return u, err
	}

	return u, nil

}

//comment update
func (commRepo *PostCommRepository) UpdateComment(u entity.Comment) error {

	_, err := commRepo.conn.Exec("UPDATE comments SET userID =$1,name=$2,message=$3,email=$4,rating =$5, 5commTime=$6 WHERE id=$7",
		u.UserID, u.Name, u.Message, u.Email, u.Rating, u.PlacedAt, u.ID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

//comment delete
func (commRepo *PostCommRepository) DeleteComment(id int) error {

	_, err := commRepo.conn.Exec("DELETE FROM comments WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

//comment store
func (commRepo *PostCommRepository) StoreComment(u entity.Comment) error {

	_, err := commRepo.conn.Exec("INSERT INTO comment (userID, name, message, email, rating, cTime)"+
		" values($1, $2, $3, $4, $5, $6)", u.UserID, u.Name, u.Message, u.Email, u.Rating, u.PlacedAt)

	if err != nil {
		//panic(err)
		return errors.New("Insertion has failed")
	}

	return nil
}
