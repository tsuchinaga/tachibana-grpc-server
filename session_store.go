package tachibana_grpc_server

import (
	"sync"
)

type iSessionStore interface {
	getSession(key string) *accountSession
	save(key string, session *accountSession)
	remove(key string)
}

type sessionStore struct {
	store map[string]*accountSession
	mtx   sync.RWMutex
}

func (s *sessionStore) getSession(key string) *accountSession {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	return s.store[key]
}

func (s *sessionStore) save(key string, session *accountSession) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.store[key] = session
}

func (s *sessionStore) remove(key string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	delete(s.store, key)
}
