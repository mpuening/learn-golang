package logger

import (
	"log"
	"os"
)

type Logger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	fatal *log.Logger
}

func NewLogger(prefix string) *Logger {
	return &Logger{
		debug: log.New(os.Stdout, "DEBUG["+prefix+"] ", log.Ldate|log.Ltime|log.Lmicroseconds),
		info:  log.New(os.Stdout, "INFO["+prefix+"] ", log.Ldate|log.Ltime|log.Lmicroseconds),
		warn:  log.New(os.Stdout, "WARN["+prefix+"] ", log.Ldate|log.Ltime|log.Lmicroseconds),
		fatal: log.New(os.Stdout, "FATAL["+prefix+"] ", log.Ldate|log.Ltime|log.Lmicroseconds),
	}
}

func (this *Logger) Debug(args ...interface{}) {
	this.debug.Println(args...)
}

func (this *Logger) Debugf(format string, args ...interface{}) {
	this.debug.Printf(format, args...)
}

func (this *Logger) Info(args ...interface{}) {
	this.info.Println(args...)
}

func (this *Logger) Infof(format string, args ...interface{}) {
	this.info.Printf(format, args...)
}

func (this *Logger) Warn(args ...interface{}) {
	this.warn.Println(args...)
}

func (this *Logger) Warnf(format string, args ...interface{}) {
	this.warn.Printf(format, args...)
}

func (this *Logger) Fatal(args ...interface{}) {
	this.fatal.Println(args...)
}

func (this *Logger) Fatalf(format string, args ...interface{}) {
	this.fatal.Printf(format, args...)
}
