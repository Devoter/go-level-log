package log

// BaseLogger is a base implementation of Logger interface.
type BaseLogger struct {
	level int
	unit  string
	name  string
}

// NewBaseLogger returns an initialized BaseLogger instance.
func NewBaseLogger(level int, unit string, name string) Logger {
	return &BaseLogger{level: level, unit: unit, name: name}
}

// Bind returns a new logger instance with the bound context.
func (l *BaseLogger) Bind(unit string, name string) Logger {
	return &BaseLogger{level: l.level, unit: unit, name: name}
}

// Level returns a log level.
func (l *BaseLogger) Level() int {
	return l.level
}

// SetLevel sets a log level.
func (l *BaseLogger) SetLevel(level int) {
	l.level = level
}

// Name returns a logged module instance name.
func (l *BaseLogger) Name() string {
	return l.name
}

// SetName updates a logged module instance name.
func (l *BaseLogger) SetName(name string) {
	l.name = name
}

// Unit returns a logged module name.
func (l *BaseLogger) Unit() string {
	return l.unit
}

// SetUnit updates a logged module name.
func (l *BaseLogger) SetUnit(unit string) {
	l.unit = unit
}

// Log outputs a log record.
func (l *BaseLogger) Log(level int, useOutStream uint8, pairs ...interface{}) {

}
