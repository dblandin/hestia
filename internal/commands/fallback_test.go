package commands

import (
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
	"testing"
)

func TestFallback(t *testing.T) {
	notifier := notifiers.Test{}

	command := Fallback{Input: types.Input{Command: "nope"}, Notifier: &notifier}
	command.Run()

	messages := notifier.Messages
	expected := "Command `nope` not found"

	if messages[0] != expected {
		t.Fatalf("Expected `%s`, but received `%s`", expected, messages[0])
	}
}
