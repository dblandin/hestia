package commands

import (
	"fmt"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
	"strings"
)

type Help struct {
	User     string
	Input    types.Input
	Notifier notifiers.Notifier
}

func (c Help) Run() {
	var message string

	if c.Input.Args != "" {
		input := types.Input{Command: c.Input.Args}
		command := Build(c.User, input, c.Notifier)

		var output []string

		if command.HelpText() != "" {
			output = append(output, command.HelpText())
		}

		if command.HelpDescription() != "" {
			output = append(output, fmt.Sprintf("  %s", command.HelpDescription()))
		}

		for _, example := range command.HelpExamples() {
			output = append(output, fmt.Sprintf("  - %s", example))
		}

		message = fmt.Sprintf("%s", strings.Join(output, "\n"))
	} else {
		message = fmt.Sprintf("%s", c.combinedHelpText())

	}

	c.Notifier.Log(message)
}

func (c Help) HelpText() string {
	return "help [command]"
}

func (c Help) HelpDescription() string {
	return "Returns helpful information"
}

func (c Help) HelpExamples() []string {
	return []string{"help", "help echo"}
}

func (c Help) combinedHelpText() string {
	var output []string
	commands := allCommands()

	for _, command := range commands {
		output = append(output, fmt.Sprintf("- %s", command.HelpText()))
	}

	return strings.Join(output, "\n")
}
