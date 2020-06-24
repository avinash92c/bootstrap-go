package foundation

import (
	"log"

	"github.com/heirko/go-contrib/logrusHelper"
	mate "github.com/heralight/logrus_mate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	//HOOKS FOR APPENDERS
	_ "github.com/heralight/logrus_mate/hooks/file"
	_ "github.com/heralight/logrus_mate/hooks/filewithformatter"
	_ "github.com/heralight/logrus_mate/hooks/graylog"
	_ "github.com/heralight/logrus_mate/hooks/logstash"
	_ "github.com/heralight/logrus_mate/hooks/syslog"
)

var (
	logInt *logrus.Logger
	logger Logger
)

func initLogging() Logger {
	log.Println("Initializing Loggers")

	logInt = logrus.New()
	conf := unmarshalConfiguration(vcfg) //UnMarshall Configuration From Viper
	logrusHelper.SetConfig(logInt, conf)

	logger = &logType{
		log: logInt,
	}
	return logger
}

func unmarshalConfiguration(viper *viper.Viper) (conf mate.LoggerConfig) {
	err := viper.UnmarshalKey("logging", &conf)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	if err = conf.Validate(); err != nil {
		panic(err)
	}
	return
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
	Fatal(args ...interface{})

	InfoF(format string, args ...interface{})
	WarnF(format string, args ...interface{})
	DebugF(format string, args ...interface{})
	ErrorF(format string, args ...interface{})
	FatalF(format string, args ...interface{})
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
func (log *logType) Fatal(args ...interface{}) {
	log.log.Fatalln(args...)
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
func (log *logType) FatalF(format string, args ...interface{}) {
	log.log.Fatalf(format, args...)
}
