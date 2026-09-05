package config

import (
	"os"
)

type Config struct {
	Port        string
	AppEnv      string
	DefaultLang string
	Domain      string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "production"
	}

	defaultLang := os.Getenv("DEFAULT_LANG")
	if defaultLang == "" {
		defaultLang = "pt"
	}

	domain := os.Getenv("DOMAIN")
	if domain == "" {
		domain = "diegocamargo.dev"
	}

	return &Config{
		Port:        port,
		AppEnv:      appEnv,
		DefaultLang: defaultLang,
		Domain:      domain,
	}
}
