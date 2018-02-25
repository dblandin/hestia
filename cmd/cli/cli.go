package main

import (
	"github.com/codeclimate/hestia/internal/commands"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/utils"
	"log"
	"os"
	"os/user"
	"regexp"
	"strings"
)

func main() {
	re := regexp.MustCompile(`(?P<command>\w+)\s?(?P<args>.*)?`)
	input := utils.ExtractInput(strings.Join(os.Args[1:], " "), re)

	log.Printf("command = %s\n", input.Command)
	log.Printf("args = %s\n", input.Args)

	notifier := notifiers.Stdout{}

	user, _ := user.Current()

	command := commands.Build(user.Username, input, notifier)
	command.Run()
}
