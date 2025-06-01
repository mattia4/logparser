package types

import (
	m "logparser/parser/models"
	"regexp"
)

const tagAndroid = "AndroidLog"
const reAndroidLog = `^(?<date>\d{2}-\d{2})\s+(?<time>\d{2}:\d{2}:\d{2}\.\d{3})\s+(?<pid>\d+)\s+(?<tid>\d+)\s+(?<level>[VDIWEA])\s+(?<tag>[^:]+):\s*(?<message>.*)$`

var reAndroid = regexp.MustCompile(reAndroidLog)

var androidLogParser = m.LogParser{
	Name:  tagAndroid,
	Regex: reAndroid,
	ParseFn: func(matches []string, rawLine string) m.LogResult {
		date, time, pid, tid, level, tag, message := matches[1], matches[2], matches[3], matches[4], matches[5], matches[6], matches[7]

		res := m.AndroidLogEntry{
			FormatTag: tagAndroid,
			Date:      date,
			Time:      time,
			Pid:       pid,
			Tid:       tid,
			Level:     level,
			Tag:       tag,
			Message:   message,
		}

		cols := []m.ColTemplate{
			{Name: "Date", Value: "Date"},
			{Name: "Time", Value: "Time"},
			{Name: "Pid", Value: "PID"},
			{Name: "Tid", Value: "TID"},
			{Name: "Level", Value: "Level"},
			{Name: "Tag", Value: "Tag"},
			{Name: "Message", Value: "Message"},
		}

		return m.LogResult{
			RawLine:    rawLine,
			FormatTag:  tagAndroid,
			ParsedData: res,
			Cols:       cols,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, androidLogParser)
}
