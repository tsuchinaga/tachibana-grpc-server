package tachibana_grpc_server

import (
	"sync"
)

type iSessionStore interface {
	getBySessionToken(token string) *accountSession
	getByClientToken(token string) *accountSession
	save(sessionToken string, clientToken string, session *accountSession)
	addClientToken(sessionToken string, clientToken string)
	removeClient(token string)
	clear()
}

type sessionStore struct {
	sessions     map[string]*accountSession
	clientTokens map[string]string
	mtx          sync.RWMutex
}

func (s *sessionStore) getBySessionToken(token string) *accountSession {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	return s.sessions[token]
}

func (s *sessionStore) getByClientToken(token string) *accountSession {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	sessionToken, ok := s.clientTokens[token]
	if !ok {
		return nil
	}

	return s.sessions[sessionToken]
}

func (s *sessionStore) save(sessionToken string, clientToken string, session *accountSession) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.sessions[sessionToken] = session
	s.clientTokens[clientToken] = sessionToken
}

func (s *sessionStore) addClientToken(sessionToken string, clientToken string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.clientTokens[clientToken] = sessionToken
}

func (s *sessionStore) removeClient(token string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	delete(s.clientTokens, token)
}

func (s *sessionStore) clear() {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.sessions = map[string]*accountSession{}
	s.clientTokens = map[string]string{}
}
