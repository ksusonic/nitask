package config

import "time"

type Config struct {
	Server  ServerConfig
	MongoDB MongoDBConfig
	Logger  LoggerConfig
}

type ServerConfig struct {
	Address string
}

type MongoDBConfig struct {
	URI            string
	MaxPoolSize    uint64
	ConnectTimeout time.Duration

	TaskDB string
}

type LoggerConfig struct {
	Level  string
	Format string
	Output string
}
