package storage

import (
	"fmt"
	"os"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/adrg/xdg"
	"github.com/dgraph-io/badger/v4"
	"github.com/labstack/gommon/log"
)

type storage struct {
	db *badger.DB
}

func NewStorageService() (StorageService, error) {
	path := fmt.Sprintf("%s/%s", xdg.DataHome, internal.AppName)

	err := os.MkdirAll(path, 0755) // Return nil if the directory already exists
	if err != nil {
		return nil, err
	}

	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}

	return &storage{
		db: db,
	}, nil
}

func (s *storage) ListStores() ([]internal.SessionData, error) {
	sessions := make([]internal.SessionData, 0)

	opts := badger.DefaultIteratorOptions

	err := s.db.View(func(txn *badger.Txn) error {
		iter := txn.NewIterator(opts)
		defer iter.Close()

		var item *badger.Item

		for iter.Rewind(); iter.Valid(); iter.Next() {
			item = iter.Item()
			item.Value(func(val []byte) error {
				session, err := decodeObj(val)
				if err != nil {
					return err
				}

				sessions = append(sessions, *session.Session)
				return nil
			})
		}

		return nil
	})

	return sessions, err
}

func (s *storage) SetStore(store *Session) error {
	err := s.db.Update(func(txn *badger.Txn) error {
		encoded, err := encodeObj(store)
		if err != nil {
			return err
		}

		err = txn.Set([]byte(store.Session.Uid), encoded)
		return err
	})

	return err
}

func (s *storage) DeleteStore(uid string) error {
	log.Fatal("not implemented DeleteSession")
	return nil
}

func (s *storage) GetStore(uid string) (*Session, error) {
	var obj *Session

	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(uid))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			obj, err = decodeObj(val)
			return err
		})

		return err
	})

	return obj, err
}
