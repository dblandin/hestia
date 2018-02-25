package commands

import (
	"fmt"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
)

type Echo struct {
	User     string
	Input    types.Input
	Notifier notifiers.Notifier
}

func (c Echo) Run() {
	message := fmt.Sprintf("<@%s>: %s", c.User, c.Input.Args)

	c.Notifier.Log(message)
}
