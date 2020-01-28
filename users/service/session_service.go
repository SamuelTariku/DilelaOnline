package service

import (
	"../../entity"
	"../../users"
)

type SessionService struct{
	sessionRepo users.SessionRepository
}

func NewSessionService(SessRepo users.SessionRepository) *SessionService{
	return &SessionService{sessionRepo:SessRepo}
}

func (s *SessionService) Session(id string) (entity.Session,error){
	sess, err := s.sessionRepo.Session(id)
	if err != nil{
		return sess, err
	}
	return sess, nil
}

func (s *SessionService) StoreSession(sess entity.Session) error {
	err := s.sessionRepo.StoreSession(sess)
	if err != nil{
		return err
	}
	return nil
}

func (s *SessionService) DeleteSession(id string) error {
	err := s.sessionRepo.DeleteSession(id)
	if err != nil{
		return err
	}
	return nil
}
