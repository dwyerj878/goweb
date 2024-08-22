package config

import "os"

type Config struct {
	ValkeyHost string
	MongoHost  string
}

func Read_env(cfg *Config) {

	cfg.ValkeyHost = os.Getenv("VALKEY_HOST")
	cfg.MongoHost = os.Getenv("MONGO_HOST")

}
