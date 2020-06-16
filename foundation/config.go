package foundation

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	vcfg   *viper.Viper
	config *configstore
)

func initConfig(configpath string) ConfigStore {
	log.Println("Initializing Config Store")

	vcfg = viper.New()
	vcfg.SetConfigName("app") //name of config file without extension
	// vcfg.AddConfigPath("./config")
	vcfg.AddConfigPath(configpath)
	vcfg.SetConfigType("yaml")
	vcfg.AddConfigPath(".")    //optionally look for config in working directory
	err := vcfg.ReadInConfig() //find and read config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}
	vcfg.WatchConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

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
	GetConfig(key string) interface{}
	/*TBD LATER
	GetConfigX(key,defaultvalue string) string
	GetConfigAsIntX(key string, defaultvalue int) int
	GetConfigAsFloatX(key string,defaultvalue float) float
	GetConfigAsBooleanX(key string,defaultvalue bool) bool
	*/
}

func (config configstore) GetConfig(key string) interface{} {
	logger.Debug("Requested Key - " + key)
	value := config.vcfg.Get(key)
	logger.Debug(fmt.Sprintf("Requested Key - %s - Value - %v", key, value))
	return value
}
