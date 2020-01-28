package handler

import "../../../entity"



type SessionHandler struct {
	active bool
	user   entity.User
}

var OldSession SessionHandler

func SessionStart(user entity.User) {
	OldSession = SessionHandler{true, user}
}

func SessionStop() {
	OldSession = SessionHandler{false, entity.User{}}
}


func GetSession() (Session *SessionHandler) {
	return Session
}

