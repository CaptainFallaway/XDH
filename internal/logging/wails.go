package logging

import "context"

type wailsLogger struct {
	ctx context.Context
}

func NewWailsLogger(ctx context.Context) Logger {
	return &wailsLogger{ctx}
}

func (wl *wailsLogger) Debug(msg any, args ...any) {
}

func (wl *wailsLogger) Info(msg any, args ...any) {

}

func (wl *wailsLogger) Warn(msg any, args ...any) {

}

func (wl *wailsLogger) Error(msg any, args ...any) {

}

func (wl *wailsLogger) Fatal(msg any, args ...any) {

}
