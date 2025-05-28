package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server *ServerConfigs
}

func LoadConfigFile() *Config {
	godotenv.Load()
	return &Config{
		Server: &ServerConfigs{
			PortServer:  getEnv("PORT_SERVER", "8080"),
			FileVersion: getEnv("FILE_VERSION", "1.0.0"),
			DevMode:     getEnvAsBool("DEV_MODE", false),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		return value == "true" || value == "1"
	}
	return fallback
}

// func getEnvAsInt(key string, fallback int) int {
// 	if value, ok := os.LookupEnv(key); ok {
// 		if intValue, err := strconv.Atoi(value); err == nil {
// 			return intValue
// 		}
// 	}
// 	return fallback
// }
