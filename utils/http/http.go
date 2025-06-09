package http

import (
	"io/fs"
	"net/http"
)

func ServeFeFile(pattern string, fs fs.FS) {
	http.Handle(pattern, http.FileServer(http.FS(fs)))
}

// start HTTP server (main thread -> blocked)
func GoServerOrError(port string, errorHandler func(err error)) {
	if err := http.ListenAndServe(port, nil); err != nil {
		errorHandler(err)
	}
}
