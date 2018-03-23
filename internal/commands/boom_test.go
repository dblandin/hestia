package commands

import (
	"github.com/codeclimate/hestia/internal/notifiers"
	"testing"
)

func TestBoom(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Boom.Run() to panic")
		}
	}()

	notifier := notifiers.Test{}

	command := Boom{Notifier: &notifier}
	command.Run()

	messages := notifier.Messages
	expected := ":boom: testing error handling"

	if messages[0] != expected {
		t.Fatalf("Expected `%s`, but received `%s`", expected, messages[0])
	}
}
