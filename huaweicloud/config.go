package huaweicloud

import "time"

type Config struct {
	BaseHost string        `default:"myhuaweicloud.com"`
	Timeout  time.Duration `default:"50s"`
}

func NewConfig() *Config {
	cfg := Config{
		BaseHost: "myhuaweicloud.com",
		Timeout:  50 * time.Second,
	}
	return &cfg
}