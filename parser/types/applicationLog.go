package types

import (
	m "logparser/parser/models"
	"regexp"
)

const tagApplication = "ApplicationLogcat"
const applicationLogcatRegex = `^(\d{2}-\d{2})\s+(\d{2}:\d{2}:\d{2}\.\d{3})\s+(\d+)-(\d+)(?:\/[^ ]+)?\s+([VDIWEF])\/([^:]+)\s*:\s*(.*)$`

var reApplication = regexp.MustCompile(applicationLogcatRegex)

var applicationLogParser = m.LogParser{
	Name:  tagApplication,
	Regex: reApplication,
	ParseFn: func(matches []string) m.LogEntry {
		date, time, level, message := matches[1], matches[2], matches[3], matches[4]
		return m.LogEntry{
			FormatTag: tagApplication,
			Date:      &date,
			Time:      &time,
			Level:     &level,
			Message:   &message,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, applicationLogParser)
}
