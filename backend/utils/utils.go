package backend

import (
	"embed"
	"io/fs"
	"logparser/logger"
	"net/http"
)

func ServeHTTPFeFile(pattern string, fs http.FileSystem) {
	http.Handle(pattern, http.FileServer(fs))
	logger.NewLogger(true).Info("Serving frontend files from %s on %s", pattern, fs)
}

func ServeFeFile(pattern string, fs fs.FS) {
	http.Handle(pattern, http.FileServer(http.FS(fs)))

	var distFs http.FileSystem
	if fs == (embed.FS{}) {
		distFs = http.Dir(".")
		logger.NewLogger(true).Info("Serving frontend files from current directory (simulated).")
	} else {
		distFs = http.FS(fs)
		logger.NewLogger(true).Info("Serving frontend files from embedded filesystem.")
	}
	http.Handle(pattern, http.FileServer(distFs))
}

func GoBrowserOrFatal(url, port string, errorHandler func(error)) {
	logger.NewLogger(true).Info("INFO: Launching browser at %s%s (simulated)", url, port)
}

func GoServerOrError(port string, errorHandler func(error)) {
	logger.NewLogger(true).Info("INFO: Starting HTTP server on port %s (simulated)", port)
}
