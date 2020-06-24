package hooks

import (
	logrus_logstash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/heralight/logrus_mate"
	"github.com/sirupsen/logrus"
)

// LogstashHookConfig configuration struct
type LogstashHookConfig struct {
	AppName          string        `json:"app_name"`
	Protocol         string        `json:"protocol"`
	Address          string        `json:"address"`
	AlwaysSentFields logrus.Fields `json:"always_sent_fields"`
	Prefix           string        `json:"prefix"`
}

func init() {
	logrus_mate.RegisterHook("logstash", NewLogstashHook)
}

// NewLogstashHook hook configuration
func NewLogstashHook(options logrus_mate.Options) (hook logrus.Hook, err error) {
	conf := LogstashHookConfig{}
	if err = options.ToObject(&conf); err != nil {
		return
	}

	hook, err = logrus_logstash.NewHookWithFieldsAndPrefix(
		conf.Protocol,
		conf.Address,
		conf.AppName,
		conf.AlwaysSentFields,
		conf.Prefix,
	)

	return
}
