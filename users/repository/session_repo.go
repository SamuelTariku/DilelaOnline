package repository

import (
	"../../entity"
	"database/sql"
	"errors"
)

type SessionPostRepo struct {
	conn *sql.DB
}

func NewSessionRepo(Conn *sql.DB) *SessionPostRepo{
	return &SessionPostRepo{conn:Conn}
}

func (s *SessionPostRepo) Session (id string) (entity.Session, error){
	rows := s.conn.QueryRow("SELECT * from session where uuid = $1", id)

	sess := entity.Session{}

	err := rows.Scan(&sess.ID,&sess.UUID, &sess.Expires, &sess.Signingkey)
	if err != nil {
		return sess, err
	}
	return sess, nil
}

func (s *SessionPostRepo) StoreSession(sess entity.Session) error{
	_, err := s.conn.Exec("INSERT into session (uuid,expires,signingkey)" + "values($1,$2,$3)", sess.UUID, sess.Expires, sess.Signingkey)
	if err != nil {
		return errors.New("failed to store")
	}
	return nil
}

func (s *SessionPostRepo) DeleteSession(id string) error {
	_, err := s.conn.Exec("DELETE FROM session WHERE id=$1", id)
	if err != nil {
		return errors.New("failed to delete")
	}
	return nil
}