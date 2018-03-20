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
	message := fmt.Sprintf("Command `%s` not found", c.Input.Command)

	c.Notifier.Log(message)
}

func (c Fallback) HelpDescription() string {
	return ""
}

func (c Fallback) HelpText() string {
	return fmt.Sprintf("Command `%s` not found", c.Input.Command)
}

func (c Fallback) HelpExamples() []string {
	return []string{}
}
