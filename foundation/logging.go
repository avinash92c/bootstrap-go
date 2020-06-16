package foundation

import (
	"log"

	"github.com/heirko/go-contrib/logrusHelper"
	"github.com/sirupsen/logrus"
)

var (
	logInt *logrus.Logger
	logger Logger
)

func initLogging() Logger {
	log.Println("Initializing Loggers")

	logInt = logrus.New()
	var c = logrusHelper.UnmarshalConfiguration(vcfg) //UnMarshall Configuration From Viper
	logrusHelper.SetConfig(logInt, c)

	logger = &logType{
		log: logInt,
	}
	return logger
}

//GetLogger to get logger instance
func GetLogger() Logger {
	return logger
}

type logType struct {
	log *logrus.Logger
}

// Logger functions
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})

	InfoF(format string, args ...interface{})
	WarnF(format string, args ...interface{})
	DebugF(format string, args ...interface{})
	ErrorF(format string, args ...interface{})
}

func (log *logType) Info(args ...interface{}) {
	log.log.Infoln(args...)
}
func (log *logType) Warn(args ...interface{}) {
	log.log.Warnln(args...)
}
func (log *logType) Debug(args ...interface{}) {
	log.log.Debugln(args...)
}
func (log *logType) Error(args ...interface{}) {
	log.log.Errorln(args...)
}

func (log *logType) InfoF(format string, args ...interface{}) {
	log.log.Infof(format, args...)
}
func (log *logType) WarnF(format string, args ...interface{}) {
	log.log.Warnf(format, args...)
}
func (log *logType) DebugF(format string, args ...interface{}) {
	log.log.Debugf(format, args...)
}
func (log *logType) ErrorF(format string, args ...interface{}) {
	log.log.Errorf(format, args...)
}
