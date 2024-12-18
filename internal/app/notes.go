package app

import (
	"github.com/CaptainFallaway/XDH/internal"
)

type IApp interface {
	NewSessionData() internal.SessionData
	OpenFileDialog() (path string)
	InitializeSession(session internal.SessionData, path string) error
	SaveSession(session internal.SessionData)
	GetSessions() (uids []string)
	GetSession(uid string) internal.SessionData
	GetGroupings(uid string, sortingMetal string) []internal.Grouping
	UpdateGrouping(internal.Grouping)
}
