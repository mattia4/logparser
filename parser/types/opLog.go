package types

import (
	m "logparser/parser/models"
	utils_date "logparser/utils/date"
	"regexp"
	"strings"
)

const tagOp = "OpLog"

const logRegexOp = `^(?P<site>\S+)\s+(?P<ip_address>\S+)\s+-\s+-\s+\[(?P<timestamp>[^\]]+)\]\s+"(?P<request_message>[^"]+)"\s+(?P<status_code>\d{3})\s+(?<size>\d+)\s*(?<full_message>.*)$`

var reOp = regexp.MustCompile(logRegexOp)

var logParserOp = m.LogParser{
	Name:  tagOp,
	Regex: reOp,
	ParseFn: func(matches []string, rawLine string) m.LogResult {
		parsedMap := make(map[string]string)
		for i, name := range reOp.SubexpNames() {
			if i != 0 && name != "" {
				parsedMap[name] = matches[i]
			}
		}

		site := parsedMap["site"]
		ipAddress := parsedMap["ip_address"]
		timestampStr := parsedMap["timestamp"]
		requestMessage := parsedMap["request_message"]
		statusCode := parsedMap["status_code"]
		sizeStr := parsedMap["size"]
		fullMessage := parsedMap["full_message"]

		parsedDate := utils_date.ParseApacheDate(timestampStr)
		parsedTime := utils_date.ParseApacheTime(timestampStr)

		accessLogEntry := m.OpLogEntry{
			RawLine:       rawLine,
			Site:          site,
			IPAddress:     ipAddress,
			Date:          parsedDate,
			Time:          parsedTime,
			RequestString: requestMessage,
			StatusCode:    statusCode,
			Size:          sizeStr,
			Message:       strings.TrimSpace(fullMessage),
		}

		cols := []m.ColTemplate{
			{Name: "Site", Value: "Site"},
			{Name: "IPAddress", Value: "IP Address"},
			{Name: "Date", Value: "Date"},
			{Name: "Time", Value: "Time"},
			{Name: "RequestString", Value: "Request"},
			{Name: "Size", Value: "Size"},
			{Name: "Message", Value: "Extra Message"},
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
