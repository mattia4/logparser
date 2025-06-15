package types

import (
	m "logparser/parser/models"
	"regexp"
)

const tagCsv = "CsvLog"
const logRegexCsv = `^(?P<timestamp>[^,]+),(?P<field1>[^,]+),(?P<field2>[^,]+),(?P<field3>[^,]+),(?P<message>.*)$`

var reCsv = regexp.MustCompile(logRegexCsv)

var logParserCsv = m.LogParser{
	Name:  tagCsv,
	Regex: reCsv,
	ParseFn: func(id int, matches []string, rawLine string) m.LogResult {
		parsedMap := make(map[string]string)
		for i, name := range reCsv.SubexpNames() {
			if i != 0 && name != "" {
				parsedMap[name] = matches[i]
			}
		}

		entry := m.CsvLogEntry{
			RawLine: rawLine,
			Field1:  parsedMap["timestamp"],
			Field2:  parsedMap["field1"],
			Field3:  parsedMap["field2"],
		}

		cols := []m.ColTemplate{
			{Name: "Field1", Value: "Timestamp"},
			{Name: "Field2", Value: "User/Source"},
			{Name: "Field3", Value: "Action"},
			{Name: "Message", Value: "Details"},
		}

		return m.LogResult{
			Id:         id,
			RawLine:    rawLine,
			FormatTag:  tagCsv,
			ParsedData: entry,
			Cols:       cols,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, logParserCsv)
}
