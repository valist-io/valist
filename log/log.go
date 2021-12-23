package log

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	warnEmoji   = "⚠️  "
	errorEmoji  = "⛔ "
	noticeEmoji = "✨ "
)

type Logger struct {
	out *log.Logger
	err *log.Logger
}

func New() *Logger {
	return &Logger{
		out: log.New(os.Stdout, "", 0),
		err: log.New(os.Stderr, "", 0),
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.out.Printf(msg, args...)
}

func (l *Logger) Notice(msg string, args ...interface{}) {
	l.out.Printf("%s%s", noticeEmoji, color.CyanString(msg, args...))
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.err.Printf("%s%s", warnEmoji, color.YellowString(msg, args...))
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.err.Printf("%s%s", errorEmoji, color.RedString(msg, args...))
}
