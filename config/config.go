package config

type Config struct {
	Host string
	Port int
}

func NewConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: 8080,
	}
}
