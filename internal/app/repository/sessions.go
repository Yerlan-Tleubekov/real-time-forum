package repository

import "sync"

type SessionTokens struct {
	sessionTokens sync.Map
}

func NewSessionTokens() *SessionTokens {

	return &SessionTokens{
		sessionTokens: sync.Map{},
	}
}
