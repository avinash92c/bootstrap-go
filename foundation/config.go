package foundation

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	vcfg   *viper.Viper
	config *configstore
)

func initConfig(configpath string) ConfigStore {
	cfgtype := os.Getenv("CONFIG_TYPE")
	cfgformat := os.Getenv("CONFIG_FORMAT")
	envprefix := os.Getenv("ENV_CONFIG_PREFIX")

	if strings.EqualFold(cfgtype, "remote") {
		return initRemoteConfig(envprefix, cfgformat)
	}
	return initFileConfig(configpath, envprefix, cfgformat)
}

//TODO REWRITE TO FIT REMOTE CONFIG & STARTUP CONFIG
func initFileConfig(configpath, envprefix, cfgformat string) ConfigStore {
	log.Println("Initializing Config Store")
	log.Println("Reading Config From ", configpath)

	vcfg = viper.New()
	vcfg.AutomaticEnv()
	vcfg.SetConfigName("app") //name of config file without extension
	vcfg.AddConfigPath(configpath)

	if len(strings.Trim(cfgformat, "")) > 0 {
		vcfg.SetConfigType(cfgformat)
	} else {
		vcfg.SetConfigType("yaml")
	}

	vcfg.AddConfigPath(".")    //optionally look for config in working directory
	err := vcfg.ReadInConfig() //find and read config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}
	vcfg.WatchConfig()

	config = &configstore{
		vcfg: vcfg,
	}
	return config
}

func initRemoteConfig(envprefix, cfgformat string) ConfigStore {
	log.Println("Initializing Config Store")

	rmtsecure := os.Getenv("REMOTE_CONFIG_SECURE")
	rmtstore := os.Getenv("REMOTE_CONFIG_STORE")
	rmthost := os.Getenv("REMOTE_CONFIG_HOST")
	rmtkey := os.Getenv("REMOTE_CONFIG_KEY")
	rmtkeyring := os.Getenv("REMOTE_CONFIG_KEYRING")

	vcfg = viper.New()
	//ENV CONFIGS
	if len(strings.Trim(envprefix, "")) > 0 {
		vcfg.SetEnvPrefix(envprefix)
	}
	vcfg.AutomaticEnv()

	if strings.EqualFold(rmtsecure, "Y") {
		vcfg.AddSecureRemoteProvider(rmtstore, rmthost, rmtkey, rmtkeyring)
	} else {
		vcfg.AddRemoteProvider(rmtstore, rmthost, rmtkey)
	}
	vcfg.SetConfigType(cfgformat) //format of config
	err := vcfg.ReadRemoteConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	err = vcfg.WatchRemoteConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	/* //TODO CONFIG CHANGE NOTIFICATION HOOKS
	vcfg.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		//INVOKE HANDLERS TO RELOAD NECESSARY
		// log.Println(vcfg.Get("database.maxopen"))
	})
	*/

	config = &configstore{
		vcfg: vcfg,
	}
	return config
}

type configstore struct {
	vcfg *viper.Viper
}

// GetConfigStore to get configuration store
func GetConfigStore() ConfigStore {
	return config
}

// ConfigStore functions
type ConfigStore interface {
	Get(key string) interface{}
	GetConfig(key string) interface{}
	GetConfigX(key string, defaultvalue interface{}) interface{}
	/*TBD LATER
	GetConfigAsIntX(key string, defaultvalue int) int
	GetConfigAsFloatX(key string,defaultvalue float) float
	GetConfigAsBooleanX(key string,defaultvalue bool) bool
	*/
}

func (config configstore) Get(key string) interface{} {
	return config.GetConfig(key)
}

func (config configstore) GetConfig(key string) interface{} {
	logger.Debug("Requested Key - " + key)
	value := config.vcfg.Get(key)
	logger.Debug(fmt.Sprintf("Requested Key - %s - Value - %v", key, value))
	return value
}

func (config configstore) GetConfigX(key string, defaultvalue interface{}) interface{} {
	logger.Debug("Requested Key - " + key)
	value := defaultvalue
	if config.vcfg.IsSet(key) {
		value = config.vcfg.Get(key)
	}
	logger.Debug(fmt.Sprintf("Requested Key - %s - Value - %v", key, value))
	return value
}
