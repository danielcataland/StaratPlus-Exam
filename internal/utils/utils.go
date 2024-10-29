package utils

import (
	"regexp"
)

func ValidateEmail() bool {

	return true
}

func MatchWord(pattern string, text string) bool {
	matched, _ := regexp.Match(pattern, []byte(text))
	if matched {
		return true
	} else {
		return false
	}
}
