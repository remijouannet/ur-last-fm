package log

import (
	"io"
	"log"
)

var (
	loggerInfo  *log.Logger
	loggerDebug *log.Logger
	loggerError *log.Logger
	loggerFatal *log.Logger
	loggerPanic *log.Logger
	debug       bool
)

func Init(dbg bool, out io.Writer) {

	if dbg {
		debug = true
	}

	loggerInfo = log.New(out, "INFO:", log.Ldate|log.Ltime)
	loggerDebug = log.New(out, "DEBUG:", log.Ldate|log.Ltime)
	loggerError = log.New(out, "ERROR:", log.Ldate|log.Ltime)
	loggerFatal = log.New(out, "FATAL:", log.Ldate|log.Ltime)
	loggerPanic = log.New(out, "PANIC:", log.Ldate|log.Ltime)
}

func Info(msg string) {
	loggerInfo.Print(msg)
}

func Debug(msg string) {
	if debug {
		loggerDebug.Print(msg)
	}
}

func Error(msg string) {
	loggerError.Print(msg)
}

func ErrorIf(msg string, err error) {
	if err != nil {
		Error(msg)
	}
}

func Fatal(msg string) {
	loggerFatal.Fatal(msg)
}

func FatalIf(msg string, err error) {
	if err != nil {
		Fatal(msg)
	}
}

func Panic(msg string) {
	loggerPanic.Panic(msg)
}

func PanicIf(err error) {
	if err != nil {
		Panic(err.Error())
	}
}
