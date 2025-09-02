package config

import (
	"fmt"
	"os"
)

type Config struct {
	ONE    string
	TWO    string
	THREE  string
	OUTPUT func()
}

func GetConfig() *Config {
	cfg := Config{
		ONE:   getEnv("ENV_ONE", "1"),
		TWO:   getEnv("ENV_TWO", "2"),
		THREE: getEnv("ENV_THREE", "3"),
	}
	cfg.OUTPUT = func() {
		fmt.Println("ONE:", cfg.ONE, "TWO:", cfg.TWO, "THREE:", cfg.THREE)
	}
	return &cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
