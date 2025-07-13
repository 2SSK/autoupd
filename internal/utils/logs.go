package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var Logger *log.Logger

func CheckLogDir() {
	if _, err := os.Stat(LogDir); os.IsNotExist(err) {
		fmt.Println("Log directory does not exist, creating:", LogDir)
		if err := os.MkdirAll(LogDir, 0755); err != nil {
			fmt.Println("Error creating log directory:", err)
			os.Exit(1)
		}
		fmt.Println("Log directory created successfully:", LogDir)
	}

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
	Logger = log.New(multi, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func WasUpdateSuccessful() bool {
	today := time.Now().Format("2006-01-02")
	logFilePath := fmt.Sprintf("%s/%s.log", LogDir, today)

	file, err := os.Open(logFilePath)
	if err != nil {
		// Log file does not exist or cannot be opened
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Update completed successfully.") {
			return true
		}
	}

	return false
}
