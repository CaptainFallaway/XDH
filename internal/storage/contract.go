package storage

import "github.com/CaptainFallaway/XDH/internal"

// StorageService is a key value storage service
// It's the key should be the uid of [internal.SessionInfo]
type StorageService interface {
	List() (sessionData []internal.SessionInfo, err error)
	Set(store *internal.Session) error
	Get(uid string) (*internal.Session, error)
	Delete(uid string) error
	Close() error
}
