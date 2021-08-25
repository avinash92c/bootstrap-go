package hooks

import (
	"github.com/heralight/logrus_mate"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LogstashHookConfig configuration struct
type FileHookConfig struct {
	FilePath   string `json:"file_path"` //"/var/log/misc.log"
	MaxSize    int    `json:"maxsize"`
	MaxBackups int    `json:"maxbackups"`
	MaxAge     int    `json:"maxage"`
}

func init() {
	logrus_mate.RegisterHook("file", NewFileHook)
}

// NewLogstashHook hook configuration
func NewFileHook(options logrus_mate.Options) (hook logrus.Hook, err error) {
	conf := FileHookConfig{}

	if err = options.ToObject(&conf); err != nil {
		return
	}

	// Set the Lumberjack logger
	lumberjackLogger := &lumberjack.Logger{
		Filename:   conf.FilePath,
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge,
		LocalTime:  true,
	}

	hook, err = NewLumberjackHook(lumberjackLogger)

	return
}
