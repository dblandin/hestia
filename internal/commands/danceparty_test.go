package commands

import (
	"github.com/codeclimate/hestia/internal/notifiers"
	"strings"
	"testing"
)

func TestDanceParty(t *testing.T) {
	notifier := notifiers.Test{}

	command := DanceParty{Notifier: &notifier}
	command.Run()

	message := notifier.Messages[0]
	emojis := strings.Split(message, " ")

	for _, expected := range command.Emojis() {
		found := false

		for _, emoji := range emojis {
			if emoji == expected {
				found = true
			}
		}

		if found {
			continue
		} else {
			t.Fatalf("Expected `%s` to include `%s`", message, expected)
		}
	}

}
