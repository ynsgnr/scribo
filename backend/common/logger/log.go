package logger

import (
	"log"
)

// Print - logs given message with log level to default logger
func Print(level LogLevel, msg string) {
	log.Printf("%s%s", level, msg)
}

// Printf - logs given message and arguments with log level to default logger
func Printf(level LogLevel, msg string, v ...interface{}) {
	log.Printf(string(level)+msg, v)
}
