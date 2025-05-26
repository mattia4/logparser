package types

import (
	m "logparser/parser/models"
	"regexp"
)

const tagCombinedApache = "CombinedApacheLog"
const logRegexCombinedApache = `^(?P<ip>\S+) \S+ (?P<user>\S+) \[(?P<timestamp>[^\]]+)\] "(?P<method>[A-Z]+) (?P<path>\S+) (?P<protocol>HTTP\/\d\.\d)" (?P<status>\d{3}) (?P<size>\d+) "(?<referrer>[^"]*)" "(?<user_agent>[^"]*)"$`

var reCombinedApache = regexp.MustCompile(logRegexCombinedApache)

var logParserCombinedApache = m.LogParser{
	Name:  tagCombinedApache,
	Regex: reCombinedApache,
	ParseFn: func(matches []string, rawLine string) m.LogResult {
		parsedMap := make(map[string]string)
		for i, name := range reCombinedApache.SubexpNames() {
			if i != 0 && name != "" {
				parsedMap[name] = matches[i]
			}
		}

		entry := m.CombinedAccessLogEntry{
			RawLine:    rawLine,
			IPAddress:  parsedMap["ip"],
			User:       parsedMap["user"],
			Timestamp:  parsedMap["timestamp"],
			Method:     parsedMap["method"],
			Path:       parsedMap["path"],
			Protocol:   parsedMap["protocol"],
			StatusCode: parsedMap["status"],
			Size:       parsedMap["size"],
			Referrer:   parsedMap["referrer"],
			UserAgent:  parsedMap["user_agent"],
		}

		cols := []m.ColTemplate{
			{Name: "IPAddress", Value: "IP"},
			{Name: "User", Value: "User"},
			{Name: "Timestamp", Value: "Timestamp"},
			{Name: "Method", Value: "Method"},
			{Name: "Path", Value: "Path"},
			{Name: "StatusCode", Value: "Status"},
			{Name: "Size", Value: "Size"},
			{Name: "Referrer", Value: "Referrer"},
			{Name: "UserAgent", Value: "User Agent"},
		}

		return m.LogResult{
			RawLine:    rawLine,
			FormatTag:  tagCombinedApache,
			ParsedData: entry,
			Cols:       cols,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, logParserCombinedApache)
}
