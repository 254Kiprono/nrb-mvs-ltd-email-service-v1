package config

import "os"

type Config struct {
	smtpUsername string
	smtpPassword string
	smtpHost     string
	smtpPort     string
}

func LoadConfig() *Config {
	return &Config{
		smtpUsername: os.Getenv("SMTP_USERNAME"),
		smtpPassword: os.Getenv("SMTP_PASSWORD"), // admin
		smtpHost:     os.Getenv("SMTP_HOST"),     // Your admin password
		smtpPort:     os.Getenv("SMTP_PORT"),     // mysql-190680-0.cloudclusters.net

	}
}
