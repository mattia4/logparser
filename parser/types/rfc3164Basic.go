package types

import (
	m "logparser/parser/models"
	"regexp"
)

const tagRfc3164Basic = "Rfc3164BasicLog"
const logRfc3164Basic = `^(?P<month>\w{3})\s+(?P<day>\d{1,2})\s+(?P<time>\d{2}:\d{2}:\d{2})\s+(?P<hostname>\S+)\s+(?P<process>[^:]+):\s*(?<message>.*)$`

var reRfc3164Basic = regexp.MustCompile(logRfc3164Basic)

var logParserRfc3164Basic = m.LogParser{
	Name:  tagRfc3164Basic,
	Regex: reRfc3164Basic,
	ParseFn: func(matches []string, rawLine string) m.LogResult {
		parsedMap := make(map[string]string)
		for i, name := range reSyslog.SubexpNames() {
			if i != 0 && name != "" {
				parsedMap[name] = matches[i]
			}
		}

		entry := m.SyslogEntry{
			RawLine:  rawLine,
			Month:    parsedMap["month"],
			Day:      parsedMap["day"],
			Time:     parsedMap["time"],
			Hostname: parsedMap["hostname"],
			Process:  parsedMap["process"],
			Message:  parsedMap["message"],
		}

		cols := []m.ColTemplate{
			{Name: "Month", Value: "Month"},
			{Name: "Day", Value: "Day"},
			{Name: "Time", Value: "Time"},
			{Name: "Hostname", Value: "Hostname"},
			{Name: "Process", Value: "Process"},
			{Name: "Message", Value: "Message"},
		}

		return m.LogResult{
			RawLine:    rawLine,
			FormatTag:  tagSyslog,
			ParsedData: entry,
			Cols:       cols,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, logParserRfc3164Basic)
}
