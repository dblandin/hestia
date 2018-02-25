package commands

import (
	"fmt"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
)

type Fallback struct {
	User     string
	Input    types.Input
	Notifier notifiers.Notifier
}

func (c Fallback) Run() {
	message := fmt.Sprintf("<@%s> command `%s` not found", c.User, c.Input.Command)

	c.Notifier.Log(message)
}
