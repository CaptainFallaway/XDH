package app

import (
	"github.com/CaptainFallaway/XDH/internal"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// CreateSession sets the GroupingId to null in the provided session
// since there are no parsed groupings for the session yet.
// In the end then stores it in the local database.
func (app *App) CreateSession(session internal.Session) string {
	err = app.Storage.Set(uid.String(), session)
	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
		return ""
	}

	return uid.String()
}

func (app *App) ListSessions() []internal.SessionInfo {
	sessions, err := app.Storage.List()

	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
		return nil
	}

	return sessions
}

func (app *App) DeleteSession(uid string) {
	err := app.Storage.Delete(uid)
	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
	}
}

func (app *App) GetSessionData(uid string) internal.SessionInfo {
	session, err := app.Storage.Get(uid)
	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
		return internal.SessionInfo{}
	}

	return *session.Session
}

func (app *App) SetSessionData(uid string, sessionData internal.SessionInfo) {
	session, err := app.Storage.Get(uid)
	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
		return
	}

	session.Session = &sessionData

	err = app.Storage.Set(uid, session)
	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
	}
}
