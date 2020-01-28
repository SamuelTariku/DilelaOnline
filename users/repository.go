package users

import "../entity"

type UserRepository interface {
	Users() ([]entity.User, error)
	Login(email string) (entity.User, error)
	UserwithID(id int) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
	StoreUser(user entity.User) error
}

type SessionRepository interface {
	Session(id string) (entity.Session, error)
	StoreSession(sess entity.Session) error
	DeleteSession(id string) error
}