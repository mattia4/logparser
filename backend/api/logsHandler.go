package backend

import (
	"encoding/json"
	c "logparser/backend/constants"
	m "logparser/backend/models"
	"logparser/logger"
	"net/http"
	"sync"
)

type LogAPIHandler struct {
	mu           sync.RWMutex
	LogData      m.LogDataResponse
	ErrorHandler func(error)
}

// constructor
func NewLogAPIHandler(initialLogData m.LogDataResponse, handlerErrorCallback func(error)) *LogAPIHandler {
	return &LogAPIHandler{
		LogData:      initialLogData,
		ErrorHandler: handlerErrorCallback,
	}
}

func (h *LogAPIHandler) GetLogsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	h.mu.RLock()
	currentLogs := h.LogData
	h.mu.RUnlock()

	if len(currentLogs.Logs) == 0 {
		http.Error(w, "No logs available. Please parse a file first.", http.StatusNoContent)
		h.ErrorHandler(m.ApiError("Attempted to serve %s with no log data available", c.LogsEP))
		return
	}

	if err := json.NewEncoder(w).Encode(currentLogs); err != nil {
		http.Error(w, "JSON serialization error", http.StatusInternalServerError)
		h.ErrorHandler(m.ApiError("JSON serialization error for %s: %w", c.LogsEP, err))
		return
	}
	logger.NewLogger(true).Info("Logs served at endpoint: %s", c.LogsEP)
}

func (h *LogAPIHandler) RegisterHandlers() {
	http.HandleFunc(c.LogsEP, h.GetLogsHandler)

	//http.HandleFunc(api.ParseLogFileEP, h.PostParseFile)

	logger.NewLogger(true).Info("Registered log handlers: %s (GET)", c.LogsEP)
}

func (h *LogAPIHandler) UpdateLog(newLogData m.LogDataResponse) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.LogData = newLogData
	logger.NewLogger(true).Info("Logs updated in memory for the API handler.")
}
