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
