package types

import (
	m "logparser/parser/models"
	"regexp"
)

const tagDBLog = "DatabaseLog"
const logRegexDBLog = `^(?P<timestamp>\d{4}-\d{2}-\d{2}\s+\d{2}:\d{2}:\d{2}\.\d{3}\s+UTC)\s+\[(?P<process_id>\d+)\]:\s+\[(?P<session_id>[^\]]+)\]\s+user=(?<user>[^,]+),db=(?<database>[^\s]+)\s+(?<level>[A-Z]+):\s+(?<message>.*)$`

var reDBLog = regexp.MustCompile(logRegexDBLog)

var logParserDBLog = m.LogParser{
	Name:  tagDBLog,
	Regex: reDBLog,
	ParseFn: func(id int, matches []string, rawLine string) m.LogResult {
		parsedMap := make(map[string]string)
		for i, name := range reDBLog.SubexpNames() {
			if i != 0 && name != "" {
				parsedMap[name] = matches[i]
			}
		}

		entry := m.DatabaseLogEntry{
			RawLine:   rawLine,
			Timestamp: parsedMap["timestamp"],
			ProcessID: parsedMap["process_id"],
			Database:  parsedMap["database"],
			User:      parsedMap["user"],
			Message:   parsedMap["message"],
		}

		cols := []m.ColTemplate{
			{Name: "Timestamp", Value: "Timestamp"},
			{Name: "ProcessID", Value: "PID"},
			{Name: "Database", Value: "Database"},
			{Name: "User", Value: "User"},
			{Name: "Message", Value: "Message"},
		}

		return m.LogResult{
			Id:         id,
			RawLine:    rawLine,
			FormatTag:  tagDBLog,
			ParsedData: entry,
			Cols:       cols,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, logParserDBLog)
}
