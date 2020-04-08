package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	ALL_L = iota
	TRACE_L
	DEBUG_L
	INFO_L
	WARN_L
	ERROR_L
	FATAL_L
	OFF_L
)

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

var (
	_out       io.Writer  // destination for output
	mu         sync.Mutex // ensures atomic writes; protects the following fields
	_level     int
	_calldepth int
)

func Debug(format string, args ...interface{}) {
	if _level <= DEBUG_L {
		Printf("DEBUG", format, args...)
	}
}
func Info(format string, args ...interface{}) {
	if _level <= INFO_L {
		Printf("INFO", format, args...)
	}
}
func Warning(format string, args ...interface{}) {
	if _level <= WARN_L {
		Printf("WARNING", format, args...)
	}
}
func Error(format string, args ...interface{}) {
	if _level <= ERROR_L {
		Printf("ERROR", format, args...)
	}
}
func Fatal(format string, args ...interface{}) {
	Printf("FATAL", format, args...)
	os.Exit(-1)
}
func SetConfig(level, calldepth int, writer io.Writer) {
	_level = level
	_calldepth = calldepth
	_out = writer
}

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is used to recover the PC and is
// provided for generality, although at the moment on all pre-defined
// paths it will be 2.
func output(level, s string) error {
	var file string
	var line int
	mu.Lock()
	defer mu.Unlock()
	var ok bool
	_, file, line, ok = runtime.Caller(_calldepth)
	if !ok {
		file = "???"
		line = 0
	}

	_, err := _out.Write([]byte(fmt.Sprintf("%s [%s] %s:%d\t%s\n", time.Now().Format("2006-01-02 15:04:05"), level, file, line, s)))
	return err
}

// Printf calls Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(level, format string, v ...interface{}) {
	output(level, fmt.Sprintf(format, v...))
}
func init() {
	SetConfig(DEBUG_L, 3, os.Stderr)
}
