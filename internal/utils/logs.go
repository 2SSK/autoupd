package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	Logger     *log.Logger
	maxLogAge  = 45 * 24 * time.Hour // 45 days
	maxLogSize = 15 * 1024 * 1024    // 15 MB
)

func CheckLogDir() {
	if _, err := os.Stat(LogDir); os.IsNotExist(err) {
		fmt.Println("Log directory does not exist, creating:", LogDir)
		if err := os.MkdirAll(LogDir, 0755); err != nil {
			fmt.Println("Error creating log directory:", err)
			os.Exit(1)
		}
		fmt.Println("Log directory created successfully:", LogDir)
	}

	// Perform log rotation and cleanup
	RotateLogs()

	today := time.Now().Format("2006-01-02")
	logFilePath := fmt.Sprintf("%s/%s.log", LogDir, today)
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		file, err := os.Create(logFilePath)
		if err != nil {
			fmt.Println("Error creating log file:", err)
			os.Exit(1)
		}
		defer file.Close()
		fmt.Println("Log file created successfully:", logFilePath)
	}
}

func SetupLogger() {
	today := time.Now().Format("2006-01-02")
	logFilePath := fmt.Sprintf("%s/%s.log", LogDir, today)

	// Open the log file for appending
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening log file: %v\n", err)
		os.Exit(1)
	}

	// Create multi-writer to write to both terminal and log file
	multi := io.MultiWriter(os.Stdout, logFile)

	// Initialize the logger
	Logger = log.New(multi, "", log.Ldate|log.Ltime)
}

func WasUpdateSuccessful() bool {
	today := time.Now().Format("2006-01-02")
	logFilePath := fmt.Sprintf("%s/%s.log", LogDir, today)

	file, err := os.Open(logFilePath)
	if err != nil {
		return false
	}
	defer file.Close()

	// Read all lines
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check from the bottom up
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]

		switch {
		case strings.Contains(line, "Update completed successfully."):
			return true
		case strings.Contains(line, "Update failed:"):
			return false
		case strings.Contains(line, "Skipping..."):
			return false
		}
	}

	return false
}

func RotateLogs() {
	var totalSize int64 = 0
	type logFile struct {
		path string
		info fs.FileInfo
	}

	var logs []logFile

	// Scan all .log files in the log directory
	err := filepath.Walk(LogDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".log" {
			return nil
		}

		// Age check: delete files older than 45 days
		if time.Since(info.ModTime()) > maxLogAge {
			if rmErr := os.Remove(path); rmErr == nil {
				fmt.Println("Deleted old log file:", path)
			}
			return nil
		}

		// Track logs for size-based cleanup
		logs = append(logs, logFile{path: path, info: info})
		totalSize += info.Size()
		return nil
	})

	if err != nil {
		fmt.Println("Error scanning log directory:", err)
		return
	}

	// Sort logs by ModTime (oldest first)
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].info.ModTime().Before(logs[j].info.ModTime())
	})

	// Delete oldest logs until total size is under 15MB
	for totalSize > int64(maxLogSize) && len(logs) > 0 {
		toDelete := logs[0]
		logs = logs[1:]

		if err := os.Remove(toDelete.path); err == nil {
			totalSize -= toDelete.info.Size()
			fmt.Println("Deleted log to maintain size limit:", toDelete.path)
		}
	}
}
