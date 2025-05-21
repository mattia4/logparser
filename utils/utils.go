package utils

import (
	"fmt"
	"logparser/models"
	"strings"
	"time"
)

func ParseLogLine(line string) *models.LogEntry {
	if pidIndex := strings.Index(line, "("); pidIndex != -1 && strings.Contains(line, ")") {
		return parseLogWithPID(line)
	} else {
		return parseLogWithoutPID(line)
	}
}

func parseLogWithPID(line string) *models.LogEntry {
	levels := []string{"D/", "I/", "W/", "I/", "E/", "F/", "V/"}
	var levelIndex int = -1
	var level string

	for _, l := range levels {
		if idx := strings.Index(line, l); idx != -1 {
			levelIndex = idx
			level = strings.TrimSpace(l)
			level = strings.ReplaceAll(level, "/", "")
			break
		}
	}

	if levelIndex == -1 {
		return nil
	}

	prefix := strings.Fields(line[:levelIndex])

	if len(prefix) < 2 {
		return nil
	}

	prefixParsed := strings.Join(prefix, " ")
	prefixParsed = strings.ReplaceAll(prefixParsed, "]", "")
	prefixParsed = strings.ReplaceAll(prefixParsed, "[", "")

	date := formatDate(prefixParsed)
	time, err := formatTime(prefixParsed)
	if err != nil {
		fmt.Println(err)
	}

	suffix := strings.TrimSpace(line[levelIndex+2:])

	processLogLevel := strings.Split(suffix, ":")
	processPart := processLogLevel[0]

	processIStart := strings.Index(processPart, "(")
	var process string = ""
	if processIStart == -1 {
		return nil
	} else {
		process = processPart[:processIStart]
	}

	tagAndMessage := strings.SplitN(suffix, ": ", 2)

	if len(tagAndMessage) < 2 {
		return nil
	}

	pidSindex := strings.Index(suffix, "(")
	pidEindex := strings.Index(suffix, ")")

	var pid string = ""
	if pidSindex == -1 {
		return nil
	} else {
		pid = suffix[pidSindex+1 : pidEindex]

	}

	return &models.LogEntry{
		Date:    date,
		Time:    time,
		PID:     pid,
		Level:   level,
		Tag:     process,
		Message: tagAndMessage[1],
	}
}

func parseLogWithoutPID(line string) *models.LogEntry {
	levels := []string{" D ", " W ", " I ", " E ", " F ", "D/", "I/", "W/", "I/", "E/", "F/", "V/"}
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
	if len(prefix) < 3 {
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

func formatDate(logLine string) string {
	parts := strings.Split(logLine, " ")
	if len(parts) < 2 {
		return ""
	}

	date := parts[0]

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ""
	}

	return parsedDate.Format("2006-01-02")
}

func formatTime(logLine string) (string, error) {
	parts := strings.Split(logLine, " ")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid log time format")
	}

	timePart := parts[1]

	parsedTime, err := time.Parse("15:04:05.000", timePart)
	if err != nil {
		return "", err
	}

	return parsedTime.Format("15:04:05:000"), nil
}
