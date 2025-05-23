package models

import (
	"regexp"
)

type LogEntry struct {
	RawLine     string
	FormatTag   string
	Date        *string
	Time        *string
	PID         *string
	TID         *string
	Level       *string
	Tag         *string
	Message     *string
	OtherFields map[string]string
}

type LogParser struct {
	Name    string
	Regex   *regexp.Regexp
	ParseFn func(matches []string) LogEntry
}

var RegisteredParsers []LogParser
