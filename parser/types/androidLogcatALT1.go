package types

import (
	m "logparser/parser/models"
	"regexp"
)

const tagAndroidAlt1 = "AndroidLogcatAlt1"
const reAndroidLogcatAlt1 = `(?P<Date>\d{2}-\d{2}) (?P<Time>\d{2}:\d{2}:\d{2}\.\d{3}) (?P<PID>\d+)-(?P<TID>\d+)/[^\s]+ (?P<Level>[A-Z])/(?P<Tag>[^:]+): (?P<Message>.+)`

var reAndroidAlt1 = regexp.MustCompile(reAndroidLogcatAlt1)

var androidLogParserAlt1 = m.LogParser{
	Name:  tagAndroidAlt1,
	Regex: reAndroidAlt1,
	ParseFn: func(matches []string, rawLine string) m.LogResult {
		date, time, pid, tid, level, tag, message := matches[1], matches[2], matches[3], matches[4], matches[5], matches[6], matches[7]

		res := m.LogEntry{
			FormatTag: tagAndroid,
			Date:      &date,
			Time:      &time,
			Pid:       &pid,
			Tid:       &tid,
			Level:     &level,
			Tag:       &tag,
			Message:   &message,
		}

		cols := []m.ColTemplate{
			{Name: "Date", Value: "Date"},
			{Name: "Time", Value: "Time"},
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
	m.RegisteredParsers = append(m.RegisteredParsers, androidLogParserAlt1)
}
