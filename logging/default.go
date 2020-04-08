package logging

import (
	"io"
	"log"
)

type defaultLogger struct {
	debug bool
	out   *log.Logger
	err   *log.Logger
}

// New instantiates a new standard logger using the Go standard library and
// given Writer.
func New(debug bool, out io.Writer, err io.Writer) Logger {
	return &defaultLogger{
		debug: debug,
		out:   log.New(out, "INFO\t", log.Ldate|log.Ltime),
		err:   log.New(err, "ERROR\t", log.Ldate|log.Ltime),
	}
}

func (l *defaultLogger) Info(msg string, data ...F) {
	print(l.out, msg, data...)
}

func (l *defaultLogger) Debug(msg string, data ...F) {
	if !l.debug {
		return
	}

	l.Info(msg, data...)
}

func (l *defaultLogger) Error(msg string, data ...F) {
	print(l.err, msg, data...)
}

func print(l *log.Logger, msg string, data ...F) {
	if len(data) == 0 {
		l.Println(msg)
		return
	}

	l.Printf("%s %v", msg, data)
}
