package commands

import (
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
)

type Boom struct {
	User     string
	Input    types.Input
	Notifier notifiers.Notifier
}

func (c Boom) Run() {
	c.Notifier.Log(":boom:")
	panic("boom: testing error handling")
}

func (c Boom) HelpText() string {
	return "boom"
}

func (c Boom) HelpDescription() string {
	return "Triggers a go panic"
}

func (c Boom) HelpExamples() []string {
	return []string{"boom"}
}
