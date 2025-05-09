package storage

import (
	"os"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/dgraph-io/badger/v4"
)

type badgerStorage struct {
	db *badger.DB
}

func NewBadgerStorage(path string) (StorageService, error) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return nil, err
	}

	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}

	return &badgerStorage{
		db: db,
	}, nil
}

func (s *badgerStorage) List() ([]internal.SessionInfo, error) {
	sessions := make([]internal.SessionInfo, 0)

	err := s.db.View(func(txn *badger.Txn) error {
		iter := txn.NewIterator(badger.DefaultIteratorOptions)
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

func (s *badgerStorage) Set(store *internal.Session) error {
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

func (s *badgerStorage) Get(uid string) (*internal.Session, error) {
	var obj *internal.Session // For broader scope

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

func (s *badgerStorage) Delete(uid string) error {
	return s.db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(uid))
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *badgerStorage) Close() error {
	return s.db.Close()
}
