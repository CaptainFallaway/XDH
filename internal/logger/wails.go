package logging

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type wailsLogger struct {
	// The wails specific context
	ctx context.Context
}

// NewWailsLogger recieves the wails specific application context and logs with the wails runtime
func NewWailsLogger(ctx context.Context) Logger {
	return &wailsLogger{ctx}
}

func (wl *wailsLogger) Debug(format string, args ...any) {
	runtime.LogDebugf(wl.ctx, format, args...)
}

func (wl *wailsLogger) Info(format string, args ...any) {
	runtime.LogInfof(wl.ctx, format, args...)
}

func (wl *wailsLogger) Warn(format string, args ...any) {
	runtime.LogWarningf(wl.ctx, format, args...)
}

func (wl *wailsLogger) Error(format string, args ...any) {
	runtime.LogErrorf(wl.ctx, format, args...)
}

func (wl *wailsLogger) Fatal(format string, args ...any) {
	runtime.LogFatalf(wl.ctx, format, args...)
}
