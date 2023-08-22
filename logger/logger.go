package logger

// runtime 包，获取行时系统的功能
import (
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
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
		infoLogger.Printf(getPrefix()+format, v...)
	}
}
func Debug(format string, v ...any) {
	if LogLevel <= DebugLevel {
		debuglogger.Printf(getPrefix()+format, v...)
	}
}

func getCallTrace() (string, int) {
	_, file, lineno, ok := runtime.Caller(3)
	if ok {
		return file, lineno
	}
	return "", 0
}
func getPrefix() string {
	file, lineno := getCallTrace()
	path := strings.Split(file, "/")
	if len(path) > 3 {
		file = strings.Join(path[len(path)-3:], "/")
	}
	return strings.Join([]string{"\t", file, "\t", strconv.Itoa(lineno), "\t"}, "")
}
