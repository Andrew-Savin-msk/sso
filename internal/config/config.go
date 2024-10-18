package config

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	Db      Db      `toml:"db_config"`
	GRPCSrv GRPCSrv `toml:"grpc_server_config"`
	App     App     `toml:"app_config"`
}

type Db struct {
	Path string `toml:"path"`
}

type GRPCSrv struct {
	Port    string        `toml:"port"`
	Timeout time.Duration `toml:"timeout"`
}

type App struct {
	TokenTtl time.Duration `toml:"token_ttl"`
	LogLevel string        `toml:"log_level"`
}

func Load() *Config {
	var cfgEnvName string = "DOCKER_CONFIG_PATH"
	if runtime.GOOS == "windows" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("unable to load file ended with error: %s", err)
		}
		cfgEnvName = "CONFIG_PATH"
	}

	cfgPath, exists := os.LookupEnv(cfgEnvName)
	if !exists || cfgEnvName == "" {
		log.Fatal("Config path variable not set")
	}

	_, err := os.Stat(cfgPath)
	if err != nil {
		log.Fatal("Unable to load config file")
	}

	var cfg Config
	_, err = toml.DecodeFile(cfgPath, &cfg)
	if err != nil {
		log.Fatalf("Unable to load data from file: %s", err)
	}
	return &cfg
}
