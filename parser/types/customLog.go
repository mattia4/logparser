package types

import (
	m "logparser/parser/models"
	"regexp"
)

const tagCustomApp = "CustomAppLog"
const logRegexCustomApp = `^\[(?P<timestamp>[^\]]+)\]\s+\[(?P<severity>[A-Z]+)\]\s+\[(?P<module>[^\]]+)\]\s+(?<message>.*)$`

var reCustomApp = regexp.MustCompile(logRegexCustomApp)

var logParserCustomApp = m.LogParser{
	Name:  tagCustomApp,
	Regex: reCustomApp,
	ParseFn: func(id int, matches []string, rawLine string) m.LogResult {
		parsedMap := make(map[string]string)
		for i, name := range reCustomApp.SubexpNames() {
			if i != 0 && name != "" {
				parsedMap[name] = matches[i]
			}
		}

		entry := m.CustomAppLogEntry{
			RawLine:   rawLine,
			Timestamp: parsedMap["timestamp"],
			Severity:  parsedMap["severity"],
			Module:    parsedMap["module"],
			Message:   parsedMap["message"],
		}

		cols := []m.ColTemplate{
			{Name: "Timestamp", Value: "Timestamp"},
			{Name: "Severity", Value: "Severity"},
			{Name: "Module", Value: "Module"},
			{Name: "Message", Value: "Message"},
		}

		return m.LogResult{
			Id:         id,
			RawLine:    rawLine,
			FormatTag:  tagCustomApp,
			ParsedData: entry,
			Cols:       cols,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, logParserCustomApp)
}
