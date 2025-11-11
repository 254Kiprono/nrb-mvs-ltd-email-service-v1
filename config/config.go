package config

import "os"

type Config struct {
	smtpUsername string
	smtpPassword string
	smtpHost     string
	smtpPort     string
	DBUsername   string
	DBPassword   string
	DBHost       string
	DBName       string
	DBPort       string
}

func LoadConfig() *Config {
	return &Config{
		smtpUsername: os.Getenv("SMTP_USERNAME"),
		smtpPassword: os.Getenv("SMTP_PASSWORD"), // admin
		smtpHost:     os.Getenv("SMTP_HOST"),     // Your admin password
		smtpPort:     os.Getenv("SMTP_PORT"),
		DBUsername:   os.Getenv("MYSQL_USER"),
		DBPassword:   os.Getenv("MYSQL_PASSWORD"),
		DBHost:       os.Getenv("DB_HOST"),
		DBName:       os.Getenv("MYSQL_DATABASE"),
		DBPort:       os.Getenv("DB_PORT"),
	}
}
