package storage

import "github.com/CaptainFallaway/XDH/internal"

type StorageService interface {
	ListSessions() (uids []string)
	InitializeSession(session *internal.Session, groupings []internal.Grouping) error
	SaveSession(session *internal.Session) error
	GetSession(uid string) (*internal.Session, error)
	DeleteSession(uid string) error

	SaveGroupings(uid string, groupings []internal.Grouping) error
	GetGroupings(uid string) ([]internal.Grouping, error)
	DeleteGroupings(uid string) error
}
