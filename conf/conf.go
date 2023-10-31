package conf

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	DBConfig DBConfig `json:"db_config"`
}

type DBConfig struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

var config = &Config{}

func loadConf() error {
	exePath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("load conf erro ,%v", err)
	}
	configPath := exePath + "/conf.json"
	conFile, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("load conf erro ,%v", err)
	}
	data, err := io.ReadAll(conFile)
	if err != nil {
		return fmt.Errorf("load conf erro ,%v", err)
	}

	if err := json.Unmarshal(data, config); err != nil {
		return fmt.Errorf("load conf erro ,%v", err)
	}
	return nil
}
func GetConfig() *Config {

	if err := loadConf(); err != nil {
		return nil
	}

	return config
}
