package foundation

import "github.com/sirupsen/logrus"

//Init all foundation dependencies
func Init(configpath string, opts LoggingOptions) (ConfigStore, Logger) {
	configstore := initConfig(configpath)
	logger := initLogging(opts)
	// InitTracer(configstore)
	// initcache()
	return configstore, logger
}

type LoggingOptions struct {
	ExtraHooks []logrus.Hook
}
