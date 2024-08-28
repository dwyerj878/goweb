package config

import (
	"log"
	"os"
	"sync"
)

type Config struct {
	ValkeyHost string
	MongoHost  string
}

var config *Config
var lock = &sync.Mutex{}

func read_env(cfg *Config) {

	cfg.ValkeyHost = os.Getenv("VALKEY_HOST")
	// TODO - we have the valkey host : get the rest of the settings from valkey
	// TODO - Error if no valkey host ?
	cfg.MongoHost = os.Getenv("MONGO_HOST")

}

func Get_config() *Config {
	if config == nil { // no config - create one
		lock.Lock()
		defer lock.Unlock()
		if config == nil { // check if config was loaded while waiting for a lock
			config = &Config{}
			log.Print("Loading Config")
			read_env(config)
		}
	}
	return config
}
