package helpers

import (
	"fmt"
	lbm "logparser/backend/models"
	m "logparser/parser/models"
)

func HandleParseDataColumnMondel(parsedEntry m.LogResult, errorHandler func(err error)) []string {
	var orderedColNames []string
	allParsedColsMap := make(map[string]lbm.ColTemplateAPI)

	switch v := parsedEntry.ParsedData.(type) {
	case m.OpLogEntry:
		orderedColNames = AddColDefinition("Site", "Sito", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("IPAddress", "IP Address", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("RequestString", "Richiesta", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Size", "Dimensione", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Date", "Data", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Time", "Ora", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Request", "Request", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Message", "Messaggio", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Size", "Size", allParsedColsMap, orderedColNames)
	case m.AccessLogEntry:
		orderedColNames = AddColDefinition("Site", "Sito", allParsedColsMap, orderedColNames)
	case m.AndroidLogEntry:
		orderedColNames = AddColDefinition("FormatTag", "Formato Tag", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Date", "Data", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Time", "Time", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Pid", "PID", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Tid", "TID", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Level", "Livello", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Tag", "Tag", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Message", "Messaggio", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("OtherFields", "Altri Campi", allParsedColsMap, orderedColNames)
	case m.SyslogEntry:
		orderedColNames = AddColDefinition("Hostname", "Hostname", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Process", "Processo", allParsedColsMap, orderedColNames)
	case m.WindowsEventEntry:
		orderedColNames = AddColDefinition("EventID", "ID Evento", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Category", "Categoria", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("EventType", "Tipo Evento", allParsedColsMap, orderedColNames)
	case m.CsvLogEntry:
		orderedColNames = AddColDefinition("Field1", "Campo 1", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Field2", "Campo 2", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Field3", "Campo 3", allParsedColsMap, orderedColNames)
	case m.CustomAppLogEntry:
		orderedColNames = AddColDefinition("Severity", "Severità", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Module", "Modulo", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("User", "Utente", allParsedColsMap, orderedColNames)
	case m.CombinedAccessLogEntry:
		orderedColNames = AddColDefinition("User", "Utente", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Method", "Metodo", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Path", "Percorso", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Protocol", "Protocollo", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Size", "Dimensione", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Referrer", "Referrer", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("UserAgent", "User Agent", allParsedColsMap, orderedColNames)
	case m.DatabaseLogEntry:
		orderedColNames = AddColDefinition("ProcessID", "ID Processo", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("Database", "Database", allParsedColsMap, orderedColNames)
		orderedColNames = AddColDefinition("User", "Utente", allParsedColsMap, orderedColNames)
	default:
		errorHandler(fmt.Errorf("unsopported column type: %T. RawLine: %s", v, orderedColNames))
	}

	return orderedColNames
}

func HandleParseDataModel(parsedEntry m.LogResult, errorHandler func(err error)) lbm.LogEntryAPI {
	apiEntry := lbm.LogEntryAPI{
		Id:      parsedEntry.Id,
		RawLine: parsedEntry.RawLine,
		LogType: parsedEntry.FormatTag,
	}

	switch v := parsedEntry.ParsedData.(type) {
	case m.OpLogEntry:
		apiEntry.Site = v.Site
		apiEntry.IPAddress = v.IPAddress
		apiEntry.Date = v.Date
		apiEntry.Time = v.Time
		apiEntry.RequestString = v.RequestString
		apiEntry.StatusCode = v.StatusCode
		apiEntry.Size = v.Size
		apiEntry.Message = v.Message
	case m.AccessLogEntry:
		apiEntry.Site = v.Site
		apiEntry.IPAddress = v.IPAddress
		apiEntry.Date = v.Date
		apiEntry.Time = v.Time
		apiEntry.Message = v.Message
		apiEntry.StatusCode = v.StatusCode
	case m.AndroidLogEntry:
		apiEntry.FormatTag = v.FormatTag
		apiEntry.Date = v.Date
		apiEntry.Time = v.Time
		apiEntry.Pid = v.Pid
		apiEntry.Tid = v.Tid
		apiEntry.Level = v.Level
		apiEntry.Tag = v.Tag
		apiEntry.Message = v.Message
		apiEntry.OtherFields = v.OtherFields
	case m.SyslogEntry:
		apiEntry.Date = fmt.Sprintf("%s %s", v.Month, v.Day)
		apiEntry.Time = v.Time
		apiEntry.Hostname = v.Hostname
		apiEntry.Process = v.Process
		apiEntry.Message = v.Message
	case m.WindowsEventEntry:
		apiEntry.Date = v.Date
		apiEntry.Time = v.Time
		apiEntry.EventID = v.EventID
		apiEntry.Category = v.Category
		apiEntry.EventType = v.EventType
		apiEntry.Message = v.Message
	case m.CsvLogEntry:
		apiEntry.Field1 = v.Field1
		apiEntry.Field2 = v.Field2
		apiEntry.Field3 = v.Field3
	case m.CustomAppLogEntry:
		apiEntry.Timestamp = v.Timestamp
		apiEntry.Severity = v.Severity
		apiEntry.Module = v.Module
		apiEntry.User = v.User
		apiEntry.Message = v.Message
	case m.CombinedAccessLogEntry:
		apiEntry.IPAddress = v.IPAddress
		apiEntry.User = v.User
		apiEntry.Timestamp = v.Timestamp
		apiEntry.Method = v.Method
		apiEntry.Path = v.Path
		apiEntry.Protocol = v.Protocol
		apiEntry.StatusCode = v.StatusCode
		apiEntry.Size = v.Size
		apiEntry.Referrer = v.Referrer
		apiEntry.UserAgent = v.UserAgent
	case m.DatabaseLogEntry:
		apiEntry.Timestamp = v.Timestamp
		apiEntry.ProcessID = v.ProcessID
		apiEntry.User = v.User
		apiEntry.Message = v.Message
	default:
		errorHandler(fmt.Errorf("unsopported log type: %T. RawLine: %s", v, apiEntry.RawLine))
	}
	return apiEntry
}

func AddColDefinition(name string, displayName string, allParsedColsMap map[string]lbm.ColTemplateAPI, orderedColNames []string) []string {
	if _, exists := allParsedColsMap[name]; !exists {
		allParsedColsMap[name] = lbm.ColTemplateAPI{Name: name, DisplayName: displayName}
		orderedColNames = append(orderedColNames, name)
	}
	return orderedColNames
}

func AddDefaultColumns(allParsedColsMap map[string]lbm.ColTemplateAPI, orderedColNames []string) []string {
	orderedColNames = AddColDefinition("LogType", "Tipo Log", allParsedColsMap, orderedColNames)
	orderedColNames = AddColDefinition("RawLine", "Raw", allParsedColsMap, orderedColNames)
	orderedColNames = AddColDefinition("StatusCode", "Status Code", allParsedColsMap, orderedColNames)
	orderedColNames = AddColDefinition("Timestamp", "Timestamp", allParsedColsMap, orderedColNames)
	return orderedColNames
}
