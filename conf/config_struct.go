package conf

type Config struct {
	Name          string `mapstructure:"name"`
	Version       string `mapstructure:"version"`
	*ServerConfig `mapstructure:"server"`
	*LogConfig    `mapstructure:"log"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}
