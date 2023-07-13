package conf

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	*DiskConfig `mapstructure:"disk"`
	*LogConfig  `mapstructure:"log"`
}

type DiskConfig struct {
	Port string `mapstructure:"port"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}
