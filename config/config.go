package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Bot      Bot      `env-required:"true" yaml:"bot"`
	Telegram Telegram `env-required:"true" yaml:"telegram"`
}

type Bot struct {
	Owner int64  `env-required:"true" yaml:"owner"`
	Token string `env-required:"true" yaml:"token"`
}

type Telegram struct {
	APIEndpoint  string `env-required:"true" yaml:"apiEndpoint"`
	FileEndpoint string `env-required:"true" yaml:"fileEndpoint"`
}

const cfgPath = "config.yml"

func Load() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig(cfgPath, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
