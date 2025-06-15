package parser

import (
	"bufio"
	"fmt"
	lbm "logparser/backend/models"
	logger "logparser/logger"
	help "logparser/parser/helpers"
	m "logparser/parser/models"
	fu "logparser/utils/file"
)

func ParseLogLine(rowIdx int, line string, parsers []m.LogParser) m.LogResult {
	for _, p := range parsers {
		matches := p.Regex.FindStringSubmatch(line)

		if len(matches) > 0 {
			return p.ParseFn(rowIdx, matches, line)
		}
	}
	return m.LogResult{Id: 0, RawLine: line, FormatTag: "Unrecognized"}
}

func ParseLogFileOrError(filePath string, registeredParser []m.LogParser, errorHandler func(err error)) ([]lbm.LogEntryAPI, []lbm.ColTemplateAPI) {
	logger := logger.NewLogger(true)

	file := fu.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var apiEntries []lbm.LogEntryAPI

	allParsedColsMap := make(map[string]lbm.ColTemplateAPI)
	var orderedColNames []string

	//orderedColNames = help.AddDefaultColumns(allParsedColsMap, orderedColNames)

	rowIdx := 0
	for scanner.Scan() {
		line := scanner.Text()

		parsedEntry := ParseLogLine(rowIdx, line, registeredParser)
		rowIdx++

		// TODO if row not recognized then print into file or put in other fields
		if parsedEntry.ParsedData == nil {
			errorHandler(fmt.Errorf("parsers have no clue of what this row is: %s", line))
			continue
		}

		apiEntry := lbm.LogEntryAPI{}

		apiEntry = help.HandleParseDataModel(parsedEntry, func(err error) {
			logger.Info("%s", err)
		})

		//additionalCols := help.HandleParseDataColumnMondel(parsedEntry, func(err error) {
		//	logger.Info("%s", err)
		//})
		//
		//fmt.Println(additionalCols)
		apiEntry.Id = parsedEntry.Id
		apiEntry.RawLine = parsedEntry.RawLine
		apiEntry.LogType = parsedEntry.FormatTag

		apiEntries = append(apiEntries, apiEntry)

		for _, col := range parsedEntry.Cols {
			orderedColNames = help.AddColDefinition(col.Name, col.Value, allParsedColsMap, orderedColNames)
		}
	}

	if err := scanner.Err(); err != nil {
		errorHandler(fmt.Errorf("error while scanning file: %w", err))
	}

	var finalCols []lbm.ColTemplateAPI
	for _, name := range orderedColNames {
		finalCols = append(finalCols, allParsedColsMap[name])
	}

	return apiEntries, finalCols
}
