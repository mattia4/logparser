package http

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"logparser/api"
	"logparser/logger"
	brw "logparser/utils/browser"
	luf "logparser/utils/filesystem"
	"net/http"
)

func ServeFeFile(pattern string, fs fs.FS) {
	http.Handle("/", http.FileServer(http.FS(fs)))
}

func HttpSetHandler(logEntries api.LogDataResponse, errorHandler func(err error)) {
	http.HandleFunc(api.LogsEP, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if logEntries.Logs == nil {
			http.Error(w, "No logs available", http.StatusInternalServerError)
			errorHandler(fmt.Errorf("try serving /api/logs with no data"))
			return
		}

		if err := json.NewEncoder(w).Encode(logEntries); err != nil {
			http.Error(w, "JSON serialization error", http.StatusInternalServerError)
			errorHandler(err)
		}

	})
}

func ServerHttpConf(frontendFiles embed.FS, logEntries api.LogDataResponse, logger *logger.Logger) {
	distFs := luf.GetFSOrError(frontendFiles, func(err error) {
		logger.Fatal("%s", err.Error())
	})

	// Serve fe static file
	ServeFeFile("/", distFs)

	HttpSetHandler(logEntries, func(err error) {
		logger.Error("Error in serving /api/logs: %v", err)
	})

	url := Endpoint + Port

	logger.Info("Server Go started on port %s", url)

	brw.GoBrowserOrFatal(url, Port, func(err error) {
		logger.Error("%s", err)
	})

	GoServerOrError(Port, func(err error) {
		logger.Error("%s", err)
	})
}

// start HTTP server (main thread -> blocked)
func GoServerOrError(port string, errorHandler func(err error)) {
	if err := http.ListenAndServe(port, nil); err != nil {
		errorHandler(err)
	}
}
