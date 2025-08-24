package config

import "github.com/BurntSushi/toml"

func LoadConfig(path string) (cfg *Config, err error) {
	_, err = toml.DecodeFile(path, &cfg)
	return
}
