package browser

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func GoBrowserOrFatal(url string, port string, errorHandler func(err error)) {
	go func() {
		if err := openBrowser(url); err != nil {
			log.Printf("cannot open brower automatically :0 : %v. open it manually :/ %s", err, url)
			errorHandler(fmt.Errorf("cannot open brower automatically :0 : %v. open it manually :/ %s", err, url))
		}
	}()
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin": // macOS
		cmd = "open"
		args = []string{url}
	default: // Linux & OTHERS
		cmd = "xdg-open"
		args = []string{url}
	}
	return exec.Command(cmd, args...).Start()
}
