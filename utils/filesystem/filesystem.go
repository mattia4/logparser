package filesystem

import (
	"io/fs"
)

func GetFSOrError(feFiles fs.FS, errorHandler func(err error)) fs.FS {
	distFs, err := fs.Sub(feFiles, "dist")
	if err != nil {
		errorHandler(err)
	}
	return distFs
}
