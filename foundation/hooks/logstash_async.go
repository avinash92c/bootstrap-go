package hooks

import (
	"encoding/json"
	"net"
	"log"

	"github.com/heralight/logrus_mate"
	"github.com/sirupsen/logrus"

	stash "github.com/avinash92c/logrus-logstash-async"
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

	conn, err := net.Dial(conf.Protocol, conf.Address)
	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	
	hook = stash.New(conn,stash.DefaultFormatter(logrus.Fields{"app_name": conf.AppName,"always_sent_fields":conf.AlwaysSentFields}))

	return hook,nil
}