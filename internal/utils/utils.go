package utils

import (
	"github.com/codeclimate/hestia/internal/types"
	"regexp"
)

func ExtractInput(text string, re *regexp.Regexp) types.Input {
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
