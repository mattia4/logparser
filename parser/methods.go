package parser

import (
	m "logparser/parser/models"
)

func ParseLogLine(line string, parsers []m.LogParser) m.LogResult {

	for _, p := range parsers {
		matches := p.Regex.FindStringSubmatch(line)

		if len(matches) > 0 {
			return p.ParseFn(matches, line)
		}

	}
	return m.LogResult{RawLine: line, FormatTag: "Unrecognized"}
}
