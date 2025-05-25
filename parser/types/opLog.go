package types

import (
	m "logparser/parser/models"
	utils_date "logparser/utils/date"
	"regexp"
)

const tagOp = "OpLog"
const logRegexOp = `^(\S+)\s+(\S+)\s+-\s+-\s+\[([^\]]+)\]\s+"([^"]+)"\s+(\d{3})`

var reOp = regexp.MustCompile(logRegexOp)

var logParserOp = m.LogParser{
	Name:  tagOp,
	Regex: reOp,
	ParseFn: func(matches []string, rawLine string) m.LogResult {
		site := matches[1]
		ipAddress := matches[2]
		timestampStr := matches[3]
		message := matches[4]
		statusCode := matches[5]

		parsedDate := utils_date.ParseApacheDate(timestampStr)
		parsedTime := utils_date.ParseApacheTime(timestampStr)

		accessLogEntry := m.AccessLogEntry{
			RawLine:    rawLine,
			Site:       site,
			IPAddress:  ipAddress,
			Date:       parsedDate,
			Time:       parsedTime,
			Message:    message,
			StatusCode: statusCode,
		}

		cols := []m.ColTemplate{
			{Name: "Site", Value: "Site"},
			{Name: "IPAddress", Value: "IP Address"},
			{Name: "Date", Value: "Date"},
			{Name: "Time", Value: "Time"},
			{Name: "Message", Value: "Message"},
			{Name: "StatusCode", Value: "StatusCode"},
		}

		return m.LogResult{
			RawLine:    rawLine,
			FormatTag:  tagOp,
			ParsedData: accessLogEntry,
			Cols:       cols,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, logParserOp)
}
