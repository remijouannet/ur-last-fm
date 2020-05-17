package log

import (
	"log"
	"os"
)

var (
	loggerInfo  *log.Logger
	loggerDebug *log.Logger
	loggerError *log.Logger
	loggerFatal *log.Logger
	loggerPanic *log.Logger
	debug       bool
)

func Init(dbg bool) {

	if dbg {
		debug = true
	}

	loggerInfo = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime)
	loggerDebug = log.New(os.Stdout, "DEBUG:", log.Ldate|log.Ltime)
	loggerError = log.New(os.Stdout, "ERROR:", log.Ldate|log.Ltime)
	loggerFatal = log.New(os.Stdout, "FATAL:", log.Ldate|log.Ltime)
	loggerPanic = log.New(os.Stdout, "PANIC:", log.Ldate|log.Ltime)
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
