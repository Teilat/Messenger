package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

type Logger struct {
	error
	Log *log.Logger
}

func NewLogger(prefix string) *Logger {
	return &Logger{
		Log: log.New(os.Stderr, prefix, log.LstdFlags),
	}
}

func (l *Logger) Info(text string) {
	fmt.Println(Blue + text + Reset)
}

func (l *Logger) Debug(text string) {
	fmt.Println(Cyan + text + Reset)
}

func (l *Logger) Warning(text string, v ...any) {
	fmt.Println(Yellow + text + Reset)
}

func (l *Logger) Error(text string, v ...any) {
	fmt.Println(Red + text + Reset)
}
