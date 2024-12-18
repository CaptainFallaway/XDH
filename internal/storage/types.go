package storage

import "github.com/CaptainFallaway/XDH/internal"

type storeObj struct {
	Session   internal.Session
	Groupings []internal.Grouping
}

func newStoreObj(session internal.Session, groupings []internal.Grouping) *storeObj {
	return &storeObj{
		Session:   session,
		Groupings: groupings,
	}
}
