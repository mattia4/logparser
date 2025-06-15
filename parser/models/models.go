package models

import (
	"regexp"
)

type OpLogEntry struct {
	RawLine       string
	Site          string
	IPAddress     string
	Date          string
	Time          string
	RequestString string
	StatusCode    string
	Size          string
	Message       string
}

type AccessLogEntry struct {
	RawLine    string
	Site       string
	IPAddress  string
	Date       string
	Time       string
	Message    string
	StatusCode string
}

type AndroidLogEntry struct {
	RawLine     string
	FormatTag   string
	Date        string
	Time        string
	Pid         string
	Tid         string
	Level       string
	Tag         string
	Message     string
	OtherFields map[string]string
}

type SyslogEntry struct {
	RawLine  string
	Month    string
	Day      string
	Time     string
	Hostname string
	Process  string
	Message  string
}

type WindowsEventEntry struct {
	RawLine   string
	Date      string
	Time      string
	EventID   string
	Category  string
	EventType string
	Message   string
}

type CsvLogEntry struct {
	RawLine string
	Field1  string
	Field2  string
	Field3  string
}

type CustomAppLogEntry struct {
	RawLine   string
	Timestamp string
	Severity  string
	Module    string
	User      string
	Message   string
}

type CombinedAccessLogEntry struct {
	RawLine    string
	IPAddress  string
	User       string
	Timestamp  string
	Method     string
	Path       string
	Protocol   string
	StatusCode string
	Size       string
	Referrer   string
	UserAgent  string
}

type DatabaseLogEntry struct {
	RawLine   string
	Timestamp string
	ProcessID string
	Database  string
	User      string
	Message   string
}

type LogResult struct {
	Id         int
	RawLine    string
	FormatTag  string
	ParsedData interface{}
	Cols       []ColTemplate
}

type ColTemplate struct {
	Name        string
	Value       string
	PdfColProps PdfColProps
}

type PdfColProps struct {
	MaxWidth  int
	MaxHeight int
}

type Result struct {
	LogResult []LogResult
	Cols      []ColTemplate
}

type LogParser struct {
	Name    string
	Regex   *regexp.Regexp
	ParseFn func(id int, matches []string, rawLine string) LogResult
}

var RegisteredParsers []LogParser
