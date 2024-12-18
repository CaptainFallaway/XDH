package storage

import "github.com/CaptainFallaway/XDH/internal"

type StorageService interface {
	ListStores() (sessionData []internal.SessionData, err error)
	SetStore(store *Session) error
	GetStore(uid string) (*Session, error)
	DeleteStore(uid string) error
}
