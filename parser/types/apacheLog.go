package types

import (
	m "logparser/parser/models"
	"regexp"
	"strings"
)

const tagApache = "ApacheLog"
const logRegexApache = `^(?P<ip>\S+) \S+ (?P<user>\S+) \[(?P<timestamp>[^\]]+)\] "(?P<method>[A-Z]+) (?P<path>\S+) (?P<protocol>HTTP\/\d\.\d)" (?P<status>\d{3}) (?P<size>\d+)$`

var reApache = regexp.MustCompile(logRegexApache)

var logParserApache = m.LogParser{
	Name:  tagApache,
	Regex: reApache,
	ParseFn: func(id int, matches []string, rawLine string) m.LogResult {
		parsedMap := make(map[string]string)
		for i, name := range reApache.SubexpNames() {
			if i != 0 && name != "" {
				parsedMap[name] = matches[i]
			}
		}

		entry := m.AccessLogEntry{
			RawLine:    rawLine,
			Site:       parsedMap["ip"],
			IPAddress:  parsedMap["ip"],
			Date:       strings.Split(parsedMap["timestamp"], ":")[0],
			Time:       strings.Split(parsedMap["timestamp"], " ")[1],
			Message:    parsedMap["method"] + " " + parsedMap["path"],
			StatusCode: parsedMap["status"],
		}

		cols := []m.ColTemplate{
			{Name: "Site", Value: "Site"},
			{Name: "IPAddress", Value: "IP Address"},
			{Name: "Date", Value: "Date"},
			{Name: "Time", Value: "Time"},
			{Name: "Message", Value: "Request"},
			{Name: "StatusCode", Value: "Status"},
		}

		return m.LogResult{
			Id:         id,
			RawLine:    rawLine,
			FormatTag:  tagApache,
			ParsedData: entry,
			Cols:       cols,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, logParserApache)
}
