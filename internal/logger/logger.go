package logger

import (
	"fmt"
	"log"
)

type MyLog struct {
	error
	*log.Logger
	fields *LogFields
}

type LogFields map[string]interface{}

func NewLogger(log *log.Logger) *MyLog {
	return &MyLog{
		Logger: log,
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

func (l *MyLog) Info() {

}

func (l *MyLog) Debug() {

}

func (l *MyLog) Error() {

}
