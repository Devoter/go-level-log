package log

// NopLogger provides a stub for logger interface.
type NopLogger struct {
	level int
	unit  string
	name  string
}

// NewNopLogger creates an instance of NopLogger.
func NewNopLogger(level int, unit string, name string) Logger {
	return &NopLogger{level: level, unit: unit, name: name}
}

// Bind returns a new logger instance with the bound context.
func (l *NopLogger) Bind(unit string, name string) Logger {
	return &NopLogger{level: l.level, unit: unit, name: name}
}

// Level returns a log level.
func (l *NopLogger) Level() int {
	return l.level
}

// SetLevel sets a log level.
func (l *NopLogger) SetLevel(level int) {
	l.level = level
}

// Name returns a logged module instance name.
func (l *NopLogger) Name() string {
	return l.name
}

// SetName updates a logged module instance name.
func (l *NopLogger) SetName(name string) {
	l.name = name
}

// Unit returns a logged module name.
func (l *NopLogger) Unit() string {
	return l.unit
}

// SetUnit updates a logged module name.
func (l *NopLogger) SetUnit(unit string) {
	l.unit = unit
}

// Log outputs a log record.
func (l *NopLogger) Log(level int, useOutStream uint8, pairs ...interface{}) {

}
