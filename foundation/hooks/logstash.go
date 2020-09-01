package hooks

import (
	"encoding/json"

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

	//Hack for Getting Map json Through and configure always_sent_fields
	if asf, err := options.String(`always_sent_fields`); err == nil {
		asfmap := make(map[string]interface{})
		err = json.Unmarshal([]byte(asf), &asfmap)
		if err == nil {
			options[`always_sent_fields`] = asfmap
		}
	}

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
