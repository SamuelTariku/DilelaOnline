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
		err = rows.Scan(&comm.ID, &comm.Name, &comm.Message, &comm.Email, &comm.PlacedAt)
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

	log.Println(u.Name)
	log.Println(u.Message)
	log.Println(u.Email)
	log.Println(u.PlacedAt)

	if err != nil {
		return u, err
	}

	return u, nil

}

//
func (commRepo *PostCommRepository) UpdateComment(u entity.Comment) error {

	_, err := commRepo.conn.Exec("UPDATE comments SET name=$1,message=$2,email=$3,commTime=$4 WHERE id=$5",
		u.Name, u.Message, u.Email, u.PlacedAt, u.ID)
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

	_, err := commRepo.conn.Exec("INSERT INTO comment (name, message, email, cTime)"+
		" values($1, $2, $3, $4)", u.Name, u.Message, u.Email, u.PlacedAt)

	if err != nil {
		//panic(err)
		return errors.New("Insertion has failed")
	}

	return nil
}
