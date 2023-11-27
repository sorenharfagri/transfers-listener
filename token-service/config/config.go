package config

import (
	"flag"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config -.
type Config struct {
	Service  `mapstructure:"service"`
	HTTP     `mapstructure:"http"`
	Log      `mapstructure:"logger"`
	Jaeger   `mapstructure:"jaeger"`
	Provider `mapstructure:"provider"`
}

// Service  -.
type Service struct {
	Name         string `mapstructure:"name" validate:"required"`
	IsProduction *bool  `mapstructure:"production" validate:"required"`
}

// HTTP -.
type HTTP struct {
	Port uint16 `mapstructure:"port" validate:"required"`
}

// Log -.
type Log struct {
	Level string `mapstructure:"log_level" validate:"required"`
}

// Jaeger  -.
type Jaeger struct {
	Host string `mapstructure:"host" env:"JAEGER_HOST" validate:"required"`
	Port uint32 `mapstructure:"port" env:"JAEGER_PORT" validate:"required"`
}

// Jaeger  -.
type Provider struct {
	Url string `mapstructure:"url" env:"PROVIDER_URL" validate:"required"`
}

func NewConfig() (*Config, error) {

	viper.SetConfigType("yaml")
	viper.SetConfigName("default")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read deafult config: %v", err)
	}

	configPtr := flag.String("config", "development", "Configuration name")
	flag.Parse()

	viper.SetConfigName(*configPtr)
	err = viper.MergeInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %v", err)
	}

	viper.AutomaticEnv()

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal config: %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("invalid or missing configuration attributes \n%v", err)
	}

	fmt.Println("App mode:", *configPtr)

	return &cfg, nil

}
