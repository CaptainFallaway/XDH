package storage

import "github.com/CaptainFallaway/XDH/internal"

type Session struct {
	Session   *internal.SessionData
	Groupings []internal.Grouping
}

func NewStore(session *internal.SessionData, groupings []internal.Grouping) *Session {
	return &Session{
		Session:   session,
		Groupings: groupings,
	}
}
