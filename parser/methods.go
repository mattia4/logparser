package parser

import (
	m "logparser/parser/models"
)

func ParseLogLine(line string, parsers []m.LogParser) m.LogEntry {

	for _, p := range parsers {
		matches := p.Regex.FindStringSubmatch(line)

		if len(matches) > 0 {
			entry := p.ParseFn(matches)
			entry.RawLine = line
			return entry
		}

	}

	return m.LogEntry{RawLine: line}

}
