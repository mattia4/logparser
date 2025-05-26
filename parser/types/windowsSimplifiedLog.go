package types

import (
	m "logparser/parser/models"
	"regexp"
)

const tagWindows = "WindowsEvent"
const logRegexWindows = `^(?P<date>\d{4}-\d{2}-\d{2})\s+(?P<time>\d{2}:\d{2}:\d{2})\s+EventID\s+(?P<event_id>\d+)\s+\((?P<category>[^)]+)\)\s+(?P<event_type>[^:]+):\s*(?P<message>.*)$`

var reWindows = regexp.MustCompile(logRegexWindows)

var logParserWindows = m.LogParser{
	Name:  tagWindows,
	Regex: reWindows,
	ParseFn: func(matches []string, rawLine string) m.LogResult {
		parsedMap := make(map[string]string)
		for i, name := range reWindows.SubexpNames() {
			if i != 0 && name != "" {
				parsedMap[name] = matches[i]
			}
		}

		entry := m.WindowsEventEntry{
			RawLine:   rawLine,
			Date:      parsedMap["date"],
			Time:      parsedMap["time"],
			EventID:   parsedMap["event_id"],
			Category:  parsedMap["category"],
			EventType: parsedMap["event_type"],
			Message:   parsedMap["message"],
		}

		cols := []m.ColTemplate{
			{Name: "Date", Value: "Date"},
			{Name: "Time", Value: "Time"},
			{Name: "EventID", Value: "Event ID"},
			{Name: "Category", Value: "Category"},
			{Name: "EventType", Value: "Event Type"},
			{Name: "Message", Value: "Message"},
		}

		return m.LogResult{
			RawLine:    rawLine,
			FormatTag:  tagWindows,
			ParsedData: entry,
			Cols:       cols,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, logParserWindows)
}
