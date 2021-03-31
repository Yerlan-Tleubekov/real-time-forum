package repository

type SessionTokens struct {
	sessionTokens map[int]string
}

func NewSessionTokens() *SessionTokens {
	return &SessionTokens{
		sessionTokens: make(map[int]string),
	}
}
