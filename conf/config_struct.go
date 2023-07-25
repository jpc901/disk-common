package conf

type Config struct {
	Name            string `mapstructure:"name"`
	Version         string `mapstructure:"version"`
	*ServerConfig   `mapstructure:"server"`
	*LogConfig      `mapstructure:"log"`
	*MySQLConfig    `mapstructure:"mysql"`
	*RedisConfig    `mapstructure:"redis"`
	*AccountConfig  `mapstructure:"account"`
	*RabbitMQConfig `mapstructure:"rabbitmq"`
	*OSSConfig      `mapstructure:"oss"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type AccountConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type RabbitMQConfig struct {
	Host              string `mapstructure:"host"`
	Port              int    `mapstructure:"port"`
	Username          string `mapstructure:"username"`
	Password          string `mapstructure:"password"`
	TransOSSQueueName string `mapstructure:"trans_oss_queue_name"`
}

type OSSConfig struct {
	OSSBucket          string `mapstructure:"oss_bucket"`
	OSSEndpoint        string `mapstructure:"oss_endpoint"`
	OSSAccessKeyId     string `mapstructure:"oss_ak"`
	OSSAccessKeySecret string `mapstructure:"oss_sk"`
}
