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
	ParseFn: func(matches []string, rawLine string) m.LogResult {
		date, time, level, message := matches[1], matches[2], matches[3], matches[4]

		res := m.LogEntry{
			FormatTag: tagAnotherApplication,
			Date:      &date,
			Time:      &time,
			Level:     &level,
			Message:   &message,
		}

		return m.LogResult{
			RawLine:    rawLine,
			FormatTag:  tagAndroid,
			ParsedData: res,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, applicationLogParser)
}
