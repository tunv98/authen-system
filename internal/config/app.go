package config

type App struct {
	Server         Server         `mapstructure:"server"`
	MySQL          MySQL          `mapstructure:"mysql"`
	Authentication Authentication `mapstructure:"authentication"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type MySQL struct {
	HostPort              string `mapstructure:"hostPort"`
	Username              string `mapstructure:"userName"`
	PassWord              string `mapstructure:"password"`
	DatabaseName          string `mapstructure:"databaseName"`
	MaxIdleConnections    int    `mapstructure:"maxIdleConnections"`
	MaxOpenConnections    int    `mapstructure:"maxOpenConnections"`
	ConnectionMaxLifetime int    `mapstructure:"connectionMaxLifetime"`
}

type Authentication struct {
	SecretKey string `mapstructure:"secretKey"`
}
