package conf

import (
	"sync"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	cfg  *AppConfig
)

// GetConfig 获取配置
func GetConfig() *AppConfig {
	once.Do(func() {
		cfg = &AppConfig{}
	})
	return cfg
}

func (cfg *AppConfig)InitConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("viper.Unmarshal failed, err:%v\n", err)
	}

	switch viper.GetString("log.level") {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	log.Infoln("---------disk config list--------")
	for _, key := range viper.AllKeys() {
		log.Infoln(key, ":", viper.Get(key))
	}
	log.Infoln("----------------------------------")

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Infoln("配置文件修改了~~~")
		if err := viper.Unmarshal(cfg); err != nil {
			log.Errorf("viper.Unmarshal failed, err:%s \n", err)
		}
	})
}