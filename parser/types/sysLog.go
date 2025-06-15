package types

import (
	m "logparser/parser/models"
	"regexp"
)

const tagSyslog = "Syslog"
const logRegexSyslog = `^(?P<month>\w{3})\s+(?P<day>\d{1,2})\s+(?P<time>\d{2}:\d{2}:\d{2})\s+(?P<hostname>\S+)\s+(?P<process>[^:]+):\s*(?P<message>.*)$`

var reSyslog = regexp.MustCompile(logRegexSyslog)

var logParserSyslog = m.LogParser{
	Name:  tagSyslog,
	Regex: reSyslog,
	ParseFn: func(id int, matches []string, rawLine string) m.LogResult {
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
			Id:         id,
			RawLine:    rawLine,
			FormatTag:  tagSyslog,
			ParsedData: entry,
			Cols:       cols,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, logParserSyslog)
}
