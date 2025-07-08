package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

const logFileName = "laptop_uptime.log"

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error getting user home directory: %v", err)
	}

	logFilePath := filepath.Join(homeDir, logFileName)

	log.Printf("Time tracker started. Logging to: %s", logFilePath)

	startTime := time.Now()
	fmt.Printf("Session started at: %s\n", startTime.Format(time.RFC1123))

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan // block until a signal is received
		endTime := time.Now()
		duration := endTime.Sub(startTime)

		logSession(logFilePath, startTime, endTime, duration)
		fmt.Printf("\nSession ended at: %s\n", endTime.Format(time.RFC1123))
		fmt.Printf("Session duration: %s\n", formatDuration(duration))
		fmt.Println("Time tracker stopped gracefully.")
		os.Exit(0)
	}()

	select {} // block until a signal causes os.Exit(0)
}

func logSession(filePath string, start, end time.Time, duration time.Duration) {
	entry := fmt.Sprintf("Session Start: %s|End: %s|Duration: %s\n", start.Format(time.RFC1123), end.Format(time.RFC1123), formatDuration(duration))

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("error opening log file %s: %v", filePath, err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(entry); err != nil {
		log.Printf("error writing to log file %s: %v", filePath, err)
	}
}

func formatDuration(d time.Duration) string {
	days := int(d.Hours() / 24)
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	parts := []string{}
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if hours > 0 || len(parts) > 0 { // include hours if non-zero or if days are present
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if minutes > 0 || len(parts) > 0 { // include minutes if non-zero or if hours/days are present
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}
	parts = append(parts, fmt.Sprintf("%ds", seconds)) // always include seconds

	return fmt.Sprintf("%s",
		parts[0]+func() string {
			if len(parts) > 1 {
				return "" + parts[1]
			}
			return ""
		}()+func() string {
			if len(parts) > 2 {
				return "" + parts[2]
			}
			return ""
		}()+func() string {
			if len(parts) > 3 {
				return "" + parts[3]
			}
			return ""
		}(),
	)
}
