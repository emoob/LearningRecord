package logger

import (
	"log"
	"os"
)

var (
	debuglogger *log.Logger
	infoLogger  *log.Logger
	logOut      *os.File
	LogLevel    int
)

const (
	DebugLevel = iota
	InfoLevel
)

func SetLevel(l int) {
	LogLevel = l
}

func SetFile(file string) {
	var err error
	logOut, err = os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	infoLogger = log.New(logOut, "[INFO]: ", log.LstdFlags)
	debuglogger = log.New(logOut, "[DeBUG]: ", log.LstdFlags)
}
func Info(format string, v ...any) {
	if LogLevel <= InfoLevel {
		infoLogger.Printf(format, v...)
	}
}
func Debug(format string, v ...any) {
	if LogLevel <= DebugLevel {
		debuglogger.Printf(format, v...)
	}
}
