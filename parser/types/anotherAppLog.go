package types

import (
	m "logparser/parser/models"
	"regexp"
)

const tagAnotherApplication = "AnotherAppLog"
const logcatRegexAnotherApplication = `^(\d{2}-\d{2})\s+(\d{2}:\d{2}:\d{2}\.\d{3})\s+(\d+)-(\d+)(?:\/[^ ]+)?\s+([VDIWEF])\/([^:]+)\s*:\s*(.*)$`

var reAnotherApplication = regexp.MustCompile(logcatRegexAnotherApplication)

var anotherApplicationLogParser = m.LogParser{
	Name:  tagAnotherApplication,
	Regex: reAnotherApplication,
	ParseFn: func(matches []string, rawLine string) m.LogResult {
		date, time, level, message := matches[1], matches[2], matches[3], matches[4]

		res := m.LogEntry{
			FormatTag: tagAnotherApplication,
			Date:      &date,
			Time:      &time,
			Level:     &level,
			Message:   &message,
		}

		cols := []m.ColTemplate{
			{Name: "Date", Value: "Date"},
			{Name: "Time", Value: "Time"},
			{Name: "Level", Value: "Level"},
			{Name: "Message", Value: "Message"},
		}

		return m.LogResult{
			RawLine:    rawLine,
			FormatTag:  tagAndroid,
			ParsedData: res,
			Cols:       cols,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, anotherApplicationLogParser)
}
