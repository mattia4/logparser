package main

import (
	"embed"
	api "logparser/backend/api"
	be_models "logparser/backend/models"
	server "logparser/backend/server"
	logger "logparser/logger"
	parser "logparser/parser"
	parser_models "logparser/parser/models"
	_ "logparser/parser/types"
	utils_cli "logparser/utils/cli"
	utils_file "logparser/utils/file"
	utils_filesystem "logparser/utils/filesystem"
	utils_http "logparser/utils/http"
)

//go:embed dist
var frontendFiles embed.FS
var globalLogData be_models.LogDataResponse

func main() {
	logger := logger.NewLogger(true)
	var parsedApiEntries []be_models.LogEntryAPI

	parsers := parser_models.RegisteredParsers

	logFilePath := utils_cli.GetInputFilePathOrError(func(err error) {
		logger.Info("%s", err.Error())
	})

	file := utils_file.Open(logFilePath)
	defer file.Close()

	logger.Info("Start parsing file log: %s", logFilePath)
	parsedApiEntries, finalCols := parser.ParseLogFileOrError(logFilePath, parsers, func(err error) {
		logger.Info("%s", err.Error())
	})

	globalLogData = be_models.LogDataResponse{
		Logs: parsedApiEntries,
		Cols: finalCols,
	}

	//httpu.ServerHttpConf(frontendFiles, globalLogData, logger)

	logHandler := api.NewLogAPIHandler(globalLogData, func(err error) {
		logger.Error("Error in LogAPIHandler: %v", err)
	})

	distFs := utils_filesystem.GetFSOrError(frontendFiles, func(err error) {
		logger.Error("%s", err.Error())
	})

	logHandler.RegisterHandlers()

	utils_http.ServeFeFile("/", distFs)

	server.StartServer()

}
