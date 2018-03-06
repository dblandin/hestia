package commands

import (
	"fmt"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
	"math/rand"
	"strings"
)

type DanceParty struct {
	User     string
	Input    types.Input
	Notifier notifiers.Notifier
}

func (c DanceParty) Run() {
	emojis := []string{
		":creepy_mario_dance:",
		":gopher_dance:",
		":mario_luigi_dance:",
		":pusheen_dance:",
		":dancing_corgi:",
	}

	rand.Shuffle(len(emojis), func(i, j int) {
		emojis[i], emojis[j] = emojis[j], emojis[i]
	})

	message := fmt.Sprintf("<@%s>: %s", c.User, strings.Join(emojis, " "))

	c.Notifier.Log(message)
}
