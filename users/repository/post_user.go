package repository

import (
	"../../entity"
	"database/sql"
	"errors"
	"log"
)

// PsqlCategoryRepository implements the
// menu.CategoryRepository interface

type PostUserRepository struct {
	conn *sql.DB
}

func NewUserPostRepo(Conn *sql.DB) *PostUserRepository {
	return &PostUserRepository{conn: Conn}
}

// Categories returns all cateogories from the database
func (userRepo *PostUserRepository) Users() ([]entity.User, error) {

	rows, err := userRepo.conn.Query("SELECT * FROM users;")
	if err != nil {
		return nil, errors.New("could not query the database")
	}
	defer rows.Close()

	usrs := []entity.User{}

	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		usrs = append(usrs, user)
	}

	return usrs, nil
}

// Category returns a category with a given id

// Category returns a category with a given id
func (userRepo *PostUserRepository) Login(email string) (entity.User, error) {

	u := entity.User{}

	row := userRepo.conn.QueryRow("SELECT * FROM users WHERE email=$1", email)
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password)

	if err != nil {
		return u, errors.New("username or Password is incorrect")
	}
	return u, nil
}

func (userRepo *PostUserRepository) UserwithID(id int) (entity.User, error) {
	row := userRepo.conn.QueryRow("SELECT * FROM users WHERE id = $1", id)

	u := entity.User{}

	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password)

	log.Println(u.FirstName)
	log.Println(u.LastName)
	log.Println(u.Email)

	//err = bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(u.Password))

	if err != nil {
		return u, err
	}

	return u, nil

}

// UpdateCategory updates a given object with a new data
func (userRepo *PostUserRepository) UpdateUser(u entity.User) error {

	_, err := userRepo.conn.Exec("UPDATE users SET firstname=$1,lastname=$2,email=$3,password=$4 WHERE id=$5",
		u.FirstName, u.LastName, u.Email, u.Password, u.ID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteCategory removes a category from a database by its id
func (userRepo *PostUserRepository) DeleteUser(id int) error {

	_, err := userRepo.conn.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreCategory stores new category information to database
func (userRepo *PostUserRepository) StoreUser(u entity.User) error {

	_, err := userRepo.conn.Exec("INSERT INTO users (firstname,lastname,email,password)"+
		" values($1, $2, $3, $4)", u.FirstName, u.LastName, u.Email, u.Password)

	if err != nil {
		//panic(err)
		return errors.New("Insertion has failed")
	}

	return nil
}
