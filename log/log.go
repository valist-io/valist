package log

import (
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	log *log.Logger
}

func New() *Logger {
	return &Logger{log.New(os.Stdout, "", 0)}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.log.Printf(msg, args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.log.Println(color.YellowString(msg, args...))
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.log.Println(color.RedString(msg, args...))
}
