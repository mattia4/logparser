package types

import (
	m "logparser/parser/models"
	"regexp"
)

const tagAndroidAlt1 = "AndroidLogcat"
const reAndroidLogcatAlt1 = `^(\d{2}-\d{2})\s+(\d{2}:\d{2}:\d{2}\.\d{3})\s+(\d+)\s+(\d+)\s+([VDIWEF])\s+([^:]+)\s*:\s*(.*)$`

var reAndroidAlt1 = regexp.MustCompile(reAndroidLogcatAlt1)

var androidLogParserAlt1 = m.LogParser{
	Name:  tagAndroidAlt1,
	Regex: reAndroidAlt1,
	ParseFn: func(matches []string) m.LogEntry {
		date, time, pid, tid, level, tag, message := matches[1], matches[2], matches[3], matches[4], matches[5], matches[6], matches[7]
		return m.LogEntry{
			FormatTag: tagAndroidAlt1,
			Date:      &date,
			Time:      &time,
			PID:       &pid,
			TID:       &tid,
			Level:     &level,
			Tag:       &tag,
			Message:   &message,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, androidLogParserAlt1)
}
