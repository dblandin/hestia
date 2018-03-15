package commands

import (
	"github.com/codeclimate/hestia/internal/commands"
	"github.com/codeclimate/hestia/internal/types"
	"testing"
)

type TestNotifier struct {
	Messages []string
}

func (n *TestNotifier) Log(message string) {
	n.Messages = append(n.Messages, message)

}

func TestRun(t *testing.T) {
	notifier := TestNotifier{}

	command := commands.Echo{User: "test", Input: types.Input{Args: "hello"}, Notifier: &notifier}
	command.Run()

	messages := notifier.Messages
	expected := "hello"

	if messages[0] != expected {
		t.Fatalf("Expected `%s`, but received `%s`", expected, messages[0])
	}
}
