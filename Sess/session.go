package Sess

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
)

func (manager *Manager) sessionID() string{
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil{
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request)(session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if(err != nil){
		panic(err)
	}
	if(cookie.Value == ""){
		sid := manager.sessionID()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path:"/"}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}
