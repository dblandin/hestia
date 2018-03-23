package commands

import (
	"github.com/codeclimate/hestia/internal/notifiers"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	notifier := notifiers.Test{}

	command := Help{Notifier: &notifier}
	command.Run()

	message := notifier.Messages[0]

	for _, command := range allCommands() {
		expected := command.HelpText()

		if !strings.Contains(message, expected) {
			t.Fatalf("Expected `%s` to include `%s`", message, expected)
		}
	}

}
