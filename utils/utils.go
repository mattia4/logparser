package utils

import (
	"fmt"
	"html/template"
	"logparser/models"
	"strings"
)

func ParseLogLine(line string) *models.LogEntry {
	levels := []string{" D ", " W ", " I ", " E ", " F "}
	var levelIndex int = -1
	var level string

	for _, l := range levels {
		if idx := strings.Index(line, l); idx != -1 {
			levelIndex = idx
			level = strings.TrimSpace(l)
			break
		}
	}

	if levelIndex == -1 {
		return nil
	}

	prefix := strings.Fields(line[:levelIndex])
	if len(prefix) < 4 {
		return nil
	}

	suffix := strings.TrimSpace(line[levelIndex+3:])
	tagAndMessage := strings.SplitN(suffix, ": ", 2)
	if len(tagAndMessage) < 2 {
		return nil
	}

	return &models.LogEntry{
		Date:    prefix[0],
		Time:    prefix[1],
		PID:     prefix[3],
		Level:   level,
		Tag:     tagAndMessage[0],
		Message: tagAndMessage[1],
	}
}

func BuildHTMLRows(entries []models.LogEntry) string {
	var builder strings.Builder
	for _, e := range entries {
		builder.WriteString("<tr>")
		builder.WriteString(fmt.Sprintf("<td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td>",
			e.Date, e.Time, e.Level, e.Tag, e.PID, template.HTMLEscapeString(e.Message)))
		builder.WriteString("</tr>\n")
	}
	return builder.String()
}
