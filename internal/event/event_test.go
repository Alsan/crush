package event

import "testing"

func TestInit(t *testing.T) {
	Init()
}

func TestSetNonInteractive(t *testing.T) {
	SetNonInteractive(true)
	SetNonInteractive(false)
}

func TestSetContinueBySessionID(t *testing.T) {
	SetContinueBySessionID(true)
}

func TestSetContinueLastSession(t *testing.T) {
	SetContinueLastSession(true)
}

func TestAlias(t *testing.T) {
	Alias("user-123")
}

func TestFlush(t *testing.T) {
	Flush()
}

func TestError(t *testing.T) {
	Error("test error", "key", "value")
	Error(nil)
}

func TestGetID(t *testing.T) {
	id := GetID()
	if id != "" {
		t.Errorf("GetID() = %q, want empty string", id)
	}
}
