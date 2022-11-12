package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// FmtLogger provides a user-friendly logger with formatting.
type FmtLogger struct {
	BaseLogger
}

// NewFmtLogger returns an initialized FmtLogger instance.
func NewFmtLogger(level int, unit string, name string) Logger {
	return &FmtLogger{BaseLogger: BaseLogger{level: level, unit: unit, name: name}}
}

// Bind returns a new logger instance with the bound context.
func (l *FmtLogger) Bind(unit string, name string) Logger {
	return &FmtLogger{BaseLogger: BaseLogger{level: l.level, unit: unit, name: name}}
}

// Log outputs a log record.
func (l *FmtLogger) Log(level int, useOutStream uint8, pairs ...interface{}) {
	if level > l.level {
		return
	}

	var builder strings.Builder

	var file string
	var line int
	var ok bool

	_, file, line, ok = runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}

	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}

	fmt.Fprintf(&builder, "file=\"%s:%d\" time=\"%s\" ", short, line, time.Now().Format(time.RFC3339))

	if l.unit != "" {
		fmt.Fprintf(&builder, "unit=\"%s\" ", l.unit)
	}

	if l.name != "" {
		fmt.Fprintf(&builder, "name=\"%s\" ", l.name)
	}

	i := 0

	for ; i < len(pairs)-3; i += 2 {
		fmt.Fprintf(&builder, "%s=\"%+v\" ", pairs[i], pairs[i+1])
	}

	if i < len(pairs)-1 {
		fmt.Fprintf(&builder, "%s=\"%+v\"", pairs[i], pairs[i+1])
	}

	fmt.Fprintln(&builder)

	s := builder.String()

	if useOutStream != 0 {
		fmt.Fprint(os.Stdout, s)
	} else {
		fmt.Fprint(os.Stderr, s)
	}
}

func (l *FmtLogger) LogIf(condition bool, level int, useOutStream uint8, pairs ...interface{}) {
	if condition {
		l.Log(level, useOutStream, pairs...)
	}
}
