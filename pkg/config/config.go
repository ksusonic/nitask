package config

import "time"

type Config struct {
	Server  ServerConfig  `toml:"server"`
	MongoDB MongoDBConfig `toml:"mongodb"`
	Logger  LoggerConfig  `toml:"logger"`
}

type ServerConfig struct {
	Address string `toml:"address"`
	Mode    string `toml:"mode"`
}

type MongoDBConfig struct {
	URI            string        `toml:"uri"`
	MaxPoolSize    uint64        `toml:"max_pool_size"`
	ConnectTimeout time.Duration `toml:"connect_timeout"`
}

type LoggerConfig struct {
	Level  string `toml:"level"`
	Format string `toml:"format"`
	Output string `toml:"output"`
}
