package types

import (
	m "logparser/parser/models"
	"regexp"
)

const tagAndroid = "AndroidLogcat"
const reAndroidLogcat = `^(\d{2}-\d{2})\s+(\d{2}:\d{2}:\d{2}\.\d{3})\s+(\d+)-(\d+)(?:\/[^ ]+)?\s+([VDIWEF])\/([^:]+)\s*:\s*(.*)$`

var reAndroid = regexp.MustCompile(reAndroidLogcat)

var androidLogParser = m.LogParser{
	Name:  tagAndroid,
	Regex: reAndroid,
	ParseFn: func(matches []string, rawLine string) m.LogResult {
		date, time, pid, tid, level, tag, message := matches[1], matches[2], matches[3], matches[4], matches[5], matches[6], matches[7]

		res := m.LogEntry{
			FormatTag: tagAndroid,
			Date:      &date,
			Time:      &time,
			PID:       &pid,
			TID:       &tid,
			Level:     &level,
			Tag:       &tag,
			Message:   &message,
		}

		return m.LogResult{
			RawLine:    rawLine,
			FormatTag:  tagAndroid,
			ParsedData: res,
		}
	},
}

func init() {
	m.RegisteredParsers = append(m.RegisteredParsers, androidLogParser)
}
