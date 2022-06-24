package tachibana_grpc_server

import (
	"sync"
)

type iSessionStore interface {
	getSession(key string) *loginSession
	save(key string, session *loginSession)
}

type sessionStore struct {
	store map[string]*loginSession
	mtx   sync.RWMutex
}

func (s *sessionStore) getSession(key string) *loginSession {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	return s.store[key]
}

func (s *sessionStore) save(key string, session *loginSession) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.store[key] = session
}
