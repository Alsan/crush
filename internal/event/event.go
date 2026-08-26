package event

import (
	"log/slog"
)

// Metrics collection disabled for personal fork.
// All functions are no-ops.

func Init() {}

func GetID() string { return "" }

func SetNonInteractive(_ bool) {}

func SetContinueBySessionID(_ bool) {}

func SetContinueLastSession(_ bool) {}

func Alias(_ string) {}

func Flush() {}

func Error(_ any, _ ...any) {}

func send(_ string, _ ...any) {
	slog.Debug("metrics disabled, skipping event")
}
