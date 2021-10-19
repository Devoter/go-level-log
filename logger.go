package log

// Logger provides a logger interface.
type Logger interface {
	Bind(unit string, name string) Logger
	Level() int
	SetLevel(level int)
	Name() string
	SetName(name string)
	Unit() string
	SetUnit(unit string)
	Log(level int, useOutStream uint8, pairs ...interface{})
}
