package logger

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	// Очищаем перед тестом
	cleanupLogs(t)
	defer cleanupLogs(t)

	logger, err := New()
	if err != nil {
		t.Fatalf("New() returned error: %v", err)
	}

	if logger == nil {
		t.Fatal("New() returned nil logger")
	}

	// Проверяем, что директория logs создана
	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		t.Errorf("logs directory was not created")
	}

	// Проверяем, что файл лога создан
	currentDate := time.Now().Format(logDateFormat)
	logFileName := filepath.Join(logsDir, "app_"+currentDate+".log")
	if _, err := os.Stat(logFileName); os.IsNotExist(err) {
		t.Errorf("log file was not created: %s", logFileName)
	}
}

func TestEnsureLogsDir(t *testing.T) {
	cleanupLogs(t)
	defer cleanupLogs(t)

	err := ensureLogsDir()
	if err != nil {
		t.Fatalf("ensureLogsDir() returned error: %v", err)
	}

	info, err := os.Stat(logsDir)
	if err != nil {
		t.Fatalf("failed to stat logs directory: %v", err)
	}

	if !info.IsDir() {
		t.Error("logs path exists but is not a directory")
	}
}

func TestCreateLogFile(t *testing.T) {
	cleanupLogs(t)
	defer cleanupLogs(t)

	if err := ensureLogsDir(); err != nil {
		t.Fatalf("failed to create logs directory: %v", err)
	}

	file, err := createLogFile()
	if err != nil {
		t.Fatalf("createLogFile() returned error: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			t.Logf("warning: failed to close log file: %v", err)
		}
	}()

	if file == nil {
		t.Fatal("createLogFile() returned nil file")
	}

	currentDate := time.Now().Format(logDateFormat)
	expectedFileName := filepath.Join(logsDir, "app_"+currentDate+".log")
	if _, err := os.Stat(expectedFileName); os.IsNotExist(err) {
		t.Errorf("log file was not created at expected path: %s", expectedFileName)
	}
}

func TestLoggerWritesToFile(t *testing.T) {
	cleanupLogs(t)
	defer cleanupLogs(t)

	logger, err := New()
	if err != nil {
		t.Fatalf("New() returned error: %v", err)
	}

	testMessage := "test log message"
	logger.Println(testMessage)

	// Читаем содержимое файла
	currentDate := time.Now().Format(logDateFormat)
	logFileName := filepath.Join(logsDir, "app_"+currentDate+".log")
	content, err := os.ReadFile(logFileName)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	// Проверяем, что сообщение записано в файл
	if len(content) == 0 {
		t.Error("log file is empty")
	}
}

// cleanupLogs удаляет директорию logs для тестов
func cleanupLogs(t *testing.T) {
	t.Helper()
	if err := os.RemoveAll(logsDir); err != nil && !os.IsNotExist(err) {
		t.Logf("warning: failed to cleanup logs directory: %v", err)
	}
}
