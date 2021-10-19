package log

// NopLogger provides a stub for logger interface.
type NopLogger struct {
	BaseLogger
}

// NewNopLogger creates an instance of NopLogger.
func NewNopLogger(level int, unit string, name string) Logger {
	return &NopLogger{BaseLogger{level: level, unit: unit, name: name}}
}

// Bind returns a new logger instance with the bound context.
func (l *NopLogger) Bind(unit string, name string) Logger {
	return &NopLogger{BaseLogger: BaseLogger{level: l.level, unit: unit, name: name}}
}
