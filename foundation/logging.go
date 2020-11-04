package foundation

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/heirko/go-contrib/logrusHelper"
	mate "github.com/heralight/logrus_mate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	//HOOKS FOR APPENDERS
	_ "github.com/heralight/logrus_mate/hooks/file"
	_ "github.com/heralight/logrus_mate/hooks/graylog"

	// _ "github.com/heralight/logrus_mate/hooks/logstash"
	_ "github.com/avinash92c/bootstrap-go/foundation/hooks"
	"github.com/avinash92c/bootstrap-go/security"
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

	podname, ok := os.LookupEnv(`CONTAINER_ID`)
	if !ok {
		rand.Seed(time.Now().UnixNano())
		//RANDOM GENERATED ID
		podname = security.RandomString(15)
	}

	logger = &logType{
		log:         logInt,
		containerid: podname,
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
	log         *logrus.Logger
	containerid string //Docker Container ID Read From Environment Variable
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

func getCallerDetails() (file string, funcname string, line int, found bool) {
	if pc, file, line, ok := runtime.Caller(3); ok {
		file = file[strings.LastIndex(file, "/")+1:]
		funcName := runtime.FuncForPC(pc).Name()
		return file, funcName, line, true
	}
	return "", "", 0, false
}

func (log *logType) getFields() logrus.Fields {
	file, funcname, line, ok := getCallerDetails()
	fields := logrus.Fields{`container`: log.containerid}
	if ok {
		fields[`src`] = fmt.Sprintf("%s:%s:%d", file, funcname, line)
	}
	return fields
}

func (log *logType) Info(args ...interface{}) {
	log.log.WithFields(log.getFields()).Infoln(args...)
}
func (log *logType) Warn(args ...interface{}) {
	log.log.WithFields(log.getFields()).Warnln(args...)
}
func (log *logType) Debug(args ...interface{}) {
	log.log.WithFields(log.getFields()).Debugln(args...)
}
func (log *logType) Error(args ...interface{}) {
	log.log.WithFields(log.getFields()).Errorln(args...)
}
func (log *logType) Fatal(args ...interface{}) {
	log.log.WithFields(log.getFields()).Fatalln(args...)
}

func (log *logType) InfoF(format string, args ...interface{}) {
	log.log.WithFields(log.getFields()).Infof(format, args...)
}
func (log *logType) WarnF(format string, args ...interface{}) {
	log.log.WithFields(log.getFields()).Warnf(format, args...)
}
func (log *logType) DebugF(format string, args ...interface{}) {
	log.log.WithFields(log.getFields()).Debugf(format, args...)
}
func (log *logType) ErrorF(format string, args ...interface{}) {
	log.log.WithFields(log.getFields()).Errorf(format, args...)
}
func (log *logType) FatalF(format string, args ...interface{}) {
	log.log.WithFields(log.getFields()).Fatalf(format, args...)
}
