package storage

import (
	"github.com/CaptainFallaway/XDH/internal"
	"github.com/dgraph-io/badger/v4"
	"github.com/labstack/gommon/log"
)

type storage struct {
	dataPath string
	db       *badger.DB
}

func NewStorageService(dataDir string) (StorageService, error) {
	db, err := badger.Open(badger.DefaultOptions(dataDir))

	if err != nil {
		return nil, err
	}

	return &storage{
		dataPath: dataDir,
		db:       db,
	}, nil
}

func (s *storage) ListSessions() []string {
	sessions := make([]string, 0)

	opts := badger.DefaultIteratorOptions
	opts.PrefetchValues = false

	// TODO: check up on this
	// This does return a error value
	s.db.View(func(txn *badger.Txn) error {
		iter := txn.NewIterator(opts)
		defer iter.Close()

		var item *badger.Item

		for iter.Rewind(); iter.Valid(); iter.Next() {
			item = iter.Item()
			sessions = append(sessions, string(item.Key()))
		}

		return nil
	})

	return sessions
}

func (s *storage) InitializeSession(session *internal.Session, groupings []internal.Grouping) error {
	obj := newStoreObj(*session, groupings)

	err := s.db.Update(func(txn *badger.Txn) error {
		encoded, err := encodeObj(obj)
		if err != nil {
			return err
		}

		err = txn.Set([]byte(session.Uid), encoded)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (s *storage) SaveSession(session *internal.Session) error {

	return nil
}

func (s *storage) DeleteSession(uid string) error {
	log.Fatal("not implemented DeleteSession")
	return nil
}

func (s *storage) GetSession(jeff string) (*internal.Session, error) {
	return &internal.Session{}
}

func (s *storage) SaveGroupings(uid string, groupings []internal.Grouping) error {
	return nil
}

func (s *storage) GetGroupings(uid string) ([]internal.Grouping, error) {
	return nil
}

func (s *storage) DeleteGroupings(uid string) error {
	return nil
}
