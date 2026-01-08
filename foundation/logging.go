package foundation

import (
	"context"
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
	// _ "github.com/heralight/logrus_mate/hooks/file"
	_ "github.com/heralight/logrus_mate/hooks/graylog"

	// _ "github.com/heralight/logrus_mate/hooks/logstash"
	_ "github.com/avinash92c/bootstrap-go/foundation/hooks"
	"github.com/avinash92c/bootstrap-go/security"
)

var (
	logInt *logrus.Logger
	logger Logger
)

const goroutineIDKey string = "routineid"

func initLogging(opts LoggingOptions) Logger {
	log.Println("Initializing Loggers")

	logInt = logrus.New()
	conf := unmarshalConfiguration(vcfg) //UnMarshall Configuration From Viper
	logrusHelper.SetConfig(logInt, conf)

	// Application-provided hooks
	for _, h := range opts.ExtraHooks {
		logInt.AddHook(h)
	}

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

// GetLogger to get logger instance
func GetLogger() Logger {
	return logger
}

type logType struct {
	log         *logrus.Logger
	containerid string //Docker Container ID Read From Environment Variable
}

// Logger functions
type Logger interface {
	// Info(args ...interface{})
	// Warn(args ...interface{})
	// Debug(args ...interface{})
	// Error(args ...interface{})
	// Fatal(args ...interface{})

	Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})

	InfoF(ctx context.Context, format string, args ...interface{})
	WarnF(ctx context.Context, format string, args ...interface{})
	DebugF(ctx context.Context, format string, args ...interface{})
	ErrorF(ctx context.Context, format string, args ...interface{})
	// FatalF(ctx context.Context,format string, args ...interface{})
}

func getCallerDetails() (file string, funcname string, line int, found bool) {
	if pc, file, line, ok := runtime.Caller(3); ok {
		file = file[strings.LastIndex(file, "/")+1:]
		funcName := runtime.FuncForPC(pc).Name()
		return file, funcName, line, true
	}
	return "", "", 0, false
}

/*
func (log *logType) getFields() logrus.Fields {
	file, funcname, line, ok := getCallerDetails()
	fields := logrus.Fields{`container`: log.containerid}
	if ok {
		fields[`src`] = fmt.Sprintf("%s:%s:%d", file, funcname, line)
	}
	return fields
}
*/

func (log *logType) getFields(ctx context.Context) logrus.Fields {
	fields := logrus.Fields{`container`: log.containerid}

	if gid, ok := ctx.Value(goroutineIDKey).(string); ok {
		fields["routineid"] = gid
	}

	file, funcname, line, ok := getCallerDetails()
	if ok {
		fields["src"] = fmt.Sprintf("%s:%s:%d", file, funcname, line)
	}

	return fields
}

func (log *logType) Info(ctx context.Context, args ...interface{}) {
	log.log.WithFields(log.getFields(ctx)).Infoln(args...)
}
func (log *logType) Warn(ctx context.Context, args ...interface{}) {
	log.log.WithFields(log.getFields(ctx)).Warnln(args...)
}
func (log *logType) Debug(ctx context.Context, args ...interface{}) {
	log.log.WithFields(log.getFields(ctx)).Debugln(args...)
}
func (log *logType) Error(ctx context.Context, args ...interface{}) {
	log.log.WithFields(log.getFields(ctx)).Errorln(args...)
}

// func (log *logType) Fatal(args ...interface{}) {
// 	log.log.WithFields(log.getFields()).Fatalln(args...)
// }

func (log *logType) InfoF(ctx context.Context, format string, args ...interface{}) {
	log.log.WithFields(log.getFields(ctx)).Infof(format, args...)
}
func (log *logType) WarnF(ctx context.Context, format string, args ...interface{}) {
	log.log.WithFields(log.getFields(ctx)).Warnf(format, args...)
}
func (log *logType) DebugF(ctx context.Context, format string, args ...interface{}) {
	log.log.WithFields(log.getFields(ctx)).Debugf(format, args...)
}
func (log *logType) ErrorF(ctx context.Context, format string, args ...interface{}) {
	log.log.WithFields(log.getFields(ctx)).Errorf(format, args...)
}

// func (log *logType) FatalF(format string, args ...interface{}) {
// 	log.log.WithFields(log.getFields()).Fatalf(format, args...)
// }
