package app

import (
	"github.com/CaptainFallaway/XDH/internal"
)

type IApp interface {
	NewSession() internal.Session
	OpenFileDialog() (path string)
	InitializeSession(session internal.Session, path string) error
	SaveSession(session internal.Session)
	GetSessions() (uids []string)
	GetSession(uid string) internal.Session
	GetGroupings(uid string, sortingMetal string) []internal.Grouping
	UpdateGrouping(internal.Grouping)
}
