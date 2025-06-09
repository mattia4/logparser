package backend

import (
	c "logparser/backend/constants"
	"logparser/logger"
	brw "logparser/utils/browser"
	httpu "logparser/utils/http"
	"net/http"
)

func StartServer() {
	url := c.Endpoint + c.Port
	logger.NewLogger(true).Info("Server Go starting on %s", url)

	brw.GoBrowserOrFatal(url, c.Port, func(err error) {
		logger.NewLogger(true).Error("Failed to open browser: %s", err)
	})

	if err := http.ListenAndServe(c.Port, nil); err != nil {
		logger.NewLogger(true).Fatal("Failed to start HTTP server: %v", err)
	}

	httpu.GoServerOrError(c.Port, func(err error) {
		logger.NewLogger(true).Error("%s", err)
	})
}
