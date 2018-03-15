package commands

import (
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
)

type Echo struct {
	User     string
	Input    types.Input
	Notifier notifiers.Notifier
}

func (c Echo) Run() {
	message := c.Input.Args

	c.Notifier.Log(message)
}

func (c Echo) HelpText() string {
	return "echo <input>"
}

func (c Echo) HelpDescription() string {
	return "Repeats back whatever you say"
}

func (c Echo) HelpExamples() []string {
	return []string{"echo hello"}
}
