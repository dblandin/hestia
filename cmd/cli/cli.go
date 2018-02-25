package main

import (
	"github.com/codeclimate/hestia/internal/commands"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
	"log"
	"os"
	"os/user"
	"regexp"
	"strings"
)

func main() {
	input := extractInput(strings.Join(os.Args[1:], " "))

	log.Printf("command = %s\n", input.Command)
	log.Printf("args = %s\n", input.Args)

	notifier := notifiers.Stdout{}

	user, _ := user.Current()

	command := commands.Build(user.Username, input, notifier)
	command.Run()
}

func extractInput(text string) types.Input {
	re := regexp.MustCompile(`(?P<command>\w+)\s?(?P<args>.*)?`)
	match := re.FindStringSubmatch(text)
	captures := extractCaptures(re, match)

	return types.Input{Command: captures["command"], Args: captures["args"]}
}

func extractCaptures(re *regexp.Regexp, match []string) map[string]string {
	captures := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i > 0 && i <= len(match) {
			captures[name] = match[i]
		}
	}

	return captures
}
