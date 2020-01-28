package Sess

import "sync"

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

type Provider interface {
	SessionInit(ssid string)(Session, error)
	SessionRead(ssid string)(Session, error)
	SessionDestroy(ssid string)error
	SessionGC(maxLifeTime int64)
}

type Manager struct{
	cookieName string
	lock 	sync.Mutex
	provider	Provider
	maxlifetime int64

}

var provides = make(map[string]Provider)

func NewManager(provideName, cookieName string, maxlife int64)(*Manager, error){
	provider, err := provides[provideName]
	if !err {
		panic(err)
	}
	return &Manager{provider:provider, cookieName:cookieName, maxlifetime:maxlife}, nil
}

func Register(name string, provider Provider){
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup{
		panic("session: Register called twice for provider")
	}

	provides[name] = provider
}
