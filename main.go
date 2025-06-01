package main

import (
	"embed"
	api "logparser/api"
	logger "logparser/logger"
	m "logparser/parser"
	lpm "logparser/parser/models"
	_ "logparser/parser/types"
	cli "logparser/utils/cli"
	fu "logparser/utils/file"
	httpu "logparser/utils/http"
)

//go:embed dist
var frontendFiles embed.FS
var globalLogData api.LogDataResponse

func main() {
	logger := logger.NewLogger(true)

	var parsedApiEntries []api.LogEntryAPI

	parsers := lpm.RegisteredParsers

	logFilePath := cli.GetInputFilePathOrError(func(err error) {
		logger.Info("%s", err.Error())
	})

	file := fu.Open(logFilePath)
	defer file.Close()

	logger.Info("Start parsing file log: %s", logFilePath)
	parsedApiEntries, finalCols := m.ParseLogFileOrError(logFilePath, parsers, func(err error) {
		logger.Info("%s", err.Error())
	})

	globalLogData = api.LogDataResponse{
		Logs: parsedApiEntries,
		Cols: finalCols,
	}

	httpu.ServerHttpConf(frontendFiles, globalLogData, logger)
}
