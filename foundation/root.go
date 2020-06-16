package foundation

//Init all foundation dependencies
func Init(configpath string) (ConfigStore, Logger) {
	configstore := initConfig(configpath)
	logger := initLogging()
	// InitTracer(configstore)
	initcache()
	return configstore, logger
}
