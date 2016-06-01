//zapwrap wraps a zap logger into the interface the Echo framework is expecting
package zapwrap

import (
	"fmt"
	"io"

	"github.com/uber-go/zap"
)

//These are the level consts from logrus, which I assume echo is using based on it being uint8s
//I could just reference these from logrus's package but that seems like it might be fragile and causes a dependency we shouldn't need.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel uint8 = iota
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

type wrappedLogger struct {
	zap zap.Logger
}

//SetOutput panics because Zap only lets you specify output at startup time
func (wl wrappedLogger) SetOutput(w io.Writer) {
	wl.zap.Panic("Zap can only have its output set at creation time with the Output() option. SetOutput() does not work.")
}

//SetLevel converts a logrus log level into a zap log level
func (wl wrappedLogger) SetLevel(level uint8) {
	var l zap.Level
	switch level {
	case PanicLevel:
		l = zap.Panic
	case FatalLevel:
		l = zap.Fatal
	case ErrorLevel:
		l = zap.Error
	case InfoLevel:
		l = zap.Info
	case DebugLevel:
		l = zap.Debug
	}

	//TODO: zap supports the levels 'All' or 'None' but logrus doesn't.. should we figure out a way to support that?
	wl.zap.SetLevel(l)
}

//Print is converted to an Info call because Zap does not have Print()
func (wl wrappedLogger) Print(i ...interface{}) {
	wl.Info(i)
}

//Printf is converted to an Infof call because Zap does not have Printf()
func (wl wrappedLogger) Printf(s string, i ...interface{}) {
	wl.Infof(s, i...)
}

func (wl wrappedLogger) Debug(i ...interface{}) {
	wl.zap.Debug(fmt.Sprint(i...))
}
func (wl wrappedLogger) Debugf(s string, i ...interface{}) {
	wl.zap.Debug(fmt.Sprintf(s, i...))
}
func (wl wrappedLogger) Info(i ...interface{}) {
	wl.zap.Info(fmt.Sprint(i...))
}
func (wl wrappedLogger) Infof(s string, i ...interface{}) {
	wl.zap.Info(fmt.Sprintf(s, i...))
}
func (wl wrappedLogger) Warn(i ...interface{}) {
	wl.zap.Warn(fmt.Sprint(i...))
}
func (wl wrappedLogger) Warnf(s string, i ...interface{}) {
	wl.zap.Warn(fmt.Sprintf(s, i...))
}
func (wl wrappedLogger) Error(i ...interface{}) {
	wl.zap.Error(fmt.Sprint(i...))
}
func (wl wrappedLogger) Errorf(s string, i ...interface{}) {
	wl.zap.Error(fmt.Sprintf(s, i...))
}
func (wl wrappedLogger) Fatal(i ...interface{}) {
	wl.zap.Fatal(fmt.Sprint(i...))
}
func (wl wrappedLogger) Fatalf(s string, i ...interface{}) {
	wl.zap.Fatal(fmt.Sprintf(s, i...))
}

func Wrap(l zap.Logger) wrappedLogger {
	return wrappedLogger{zap: l}
}
