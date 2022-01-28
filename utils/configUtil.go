package utils

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName     string         `json:"app_name"`
	AppModel    string         `json:"app_model"`
	AppHost     string         `json:"app_host"`
	AppPort     string         `json:"app_port"`
	Database    DatabaseConfig `json:"database"`
	RedisConfig RedisConfig    `json:"redis_config"`
}

type DatabaseConfig struct {
	Driver      string `json:"driver"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	ServiceName string `json:"service_name"`
	Charset     string `json:"charset"`
	ShowSql     bool   `json:"show_sql"`
}

type RedisConfig struct {
	Addr      string `json:"addr"`
	Port      string `json:"port"`
	Pawssword string `json:"pawssword"`
	Db        int    `json:"db"`
}

var cfg *Config = nil

func GetConfig() *Config {
	return cfg
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
