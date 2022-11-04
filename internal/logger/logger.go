package logger

import (
	"fmt"
	"log"
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

type MyLog struct {
	error
	Log    *log.Logger
	fields *LogFields
}

type LogFields map[string]interface{}

func NewLogger(log *log.Logger) *MyLog {
	return &MyLog{
		Log: log,
	}
}

func (l *MyLog) iterateFields() {
	if l.fields == nil {
		return
	}
	for s, i := range *l.fields {
		fmt.Println(s, " : ", i)
	}
}

func (l *MyLog) WithFields(fields LogFields) {
	l.fields = &fields
}

func (l *MyLog) Info(text string) {
	fmt.Println(text)
}

func (l *MyLog) Debug(text string) {
	fmt.Println(Yellow + text + Reset)

}

func (l *MyLog) Error(text string, v ...any) {
	fmt.Println(Red + text + Reset)
}
