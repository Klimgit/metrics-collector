// pkg/utils/logger.go
package utils

import (
	"log"
	"os"
)

type Logger struct {
	debug bool
}

func NewLogger(debug bool) *Logger {
	return &Logger{debug: debug}
}

func (l *Logger) Info(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	if l.debug {
		log.Printf("[DEBUG] "+format, args...)
	}
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	log.Printf("[FATAL] "+format, args...)
	os.Exit(1)
}
