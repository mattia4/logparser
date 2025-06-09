package backend

import (
	"fmt"
)

type LogEntryAPI struct {
	RawLine string `json:"RawLine"`
	LogType string `json:"LogType"`

	Date      string `json:"Date,omitempty"`
	Time      string `json:"Time,omitempty"`
	Timestamp string `json:"Timestamp,omitempty"`

	Message  string `json:"Message,omitempty"`
	Category string `json:"Category,omitempty"`
	Level    string `json:"Level,omitempty"`
	Severity string `json:"Severity,omitempty"`

	IPAddress string `json:"IPAddress,omitempty"`
	Site      string `json:"Site,omitempty"`
	Hostname  string `json:"Hostname,omitempty"`
	User      string `json:"User,omitempty"`

	RequestString string `json:"RequestString,omitempty"`
	Method        string `json:"Method,omitempty"`
	Path          string `json:"Path,omitempty"`
	Protocol      string `json:"Protocol,omitempty"`
	StatusCode    string `json:"StatusCode,omitempty"`
	Size          string `json:"Size,omitempty"`
	Referrer      string `json:"Referrer,omitempty"`
	UserAgent     string `json:"UserAgent,omitempty"`

	ProcessID string `json:"ProcessID,omitempty"`
	Process   string `json:"Process,omitempty"`
	Pid       string `json:"Pid,omitempty"`
	Tid       string `json:"Tid,omitempty"`
	EventID   string `json:"EventID,omitempty"`
	EventType string `json:"EventType,omitempty"`
	Component string `json:"Component,omitempty"`
	Module    string `json:"Module,omitempty"`

	FormatTag string `json:"FormatTag,omitempty"`
	Tag       string `json:"Tag,omitempty"`

	Field1 string `json:"Field1,omitempty"`
	Field2 string `json:"Field2,omitempty"`
	Field3 string `json:"Field3,omitempty"`

	OtherFields map[string]string `json:"OtherFields,omitempty"`
}

type ColTemplateAPI struct {
	Name        string `json:"Name"`
	DisplayName string `json:"DisplayName"`
}

type LogDataResponse struct {
	Logs []LogEntryAPI    `json:"logs"`
	Cols []ColTemplateAPI `json:"cols"`
}

func ApiError(format string, a ...any) error { return fmt.Errorf(format, a...) }
