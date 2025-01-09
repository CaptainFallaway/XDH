package app

import (
	"github.com/CaptainFallaway/XDH/internal"
)

type IApp interface {
	NewSessionData() internal.SessionInfo
	OpenFileDialog() (path string)
	CreateSession(session internal.SessionInfo, path string) error
	SaveSession(session internal.SessionInfo)
	GetSessions() (uids []string)
	GetSession(uid string) internal.SessionInfo
	GetGroupings(uid string, sortingMetal string) []internal.Grouping
	UpdateGrouping(internal.Grouping)
}
