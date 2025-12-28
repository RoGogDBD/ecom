package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	logsDir        = "logs"
	logFilePattern = "app_%s.log"
	logDateFormat  = "2006-01-02"
	logFileMode    = 0644
	logsDirMode    = 0755
	logFlags       = log.LstdFlags | log.Lshortfile
	errCreateDir   = "failed to create logs directory: %w"
	errCreateFile  = "failed to create log file: %w"
)

// New создает и возвращает новый логгер, который пишет одновременно в stdout и в файл.
// Файл логов создается в директории logs с именем app_YYYY-MM-DD.log.
func New() (*log.Logger, error) {
	if err := ensureLogsDir(); err != nil {
		return nil, err
	}

	logFile, err := createLogFile()
	if err != nil {
		return nil, err
	}

	// Создаем MultiWriter для записи одновременно в stdout и файл
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger := log.New(multiWriter, "", logFlags)

	return logger, nil
}

// ensureLogsDir создает директорию logs, если она не существует.
func ensureLogsDir() error {
	if err := os.MkdirAll(logsDir, logsDirMode); err != nil {
		return fmt.Errorf(errCreateDir, err)
	}
	return nil
}

// createLogFile создает файл для логов с именем, содержащим текущую дату.
func createLogFile() (*os.File, error) {
	currentDate := time.Now().Format(logDateFormat)
	logFileName := fmt.Sprintf(logFilePattern, currentDate)
	logFilePath := filepath.Join(logsDir, logFileName)

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, logFileMode)
	if err != nil {
		return nil, fmt.Errorf(errCreateFile, err)
	}

	return file, nil
}
