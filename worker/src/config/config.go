package config

import (
    "fmt"
	"encoding/json"
	"io/ioutil"
)

type Config struct {

    DB     *DBConfig      `json:"db_config"`
	ENV    *ENVConfig     `json:"env_config"`
	Server *ServerConfig  `json:"server_config"`
}

type DBConfig struct {
	Dialect  string   `json:"dialect"`
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Charset  string   `json:"charset"`
}

type ENVConfig struct {
	Language  string  `json:"language"`
	Library   string  `json:"library"`
}

type ServerConfig struct {
	Port  string  `json:"port"`
}

func NewConfig (configPath string) *Config {

	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error opening config file.")
        panic(err.Error())
    }

	config := Config{}
	err = json.Unmarshal([]byte(file), &config)

	if err != nil {
		fmt.Println("Couldn't read config file.")
        panic(err)
	}
	fmt.Println("Config file read.")

	return &config
}

func (config *Config) GetDBConfig() *DBConfig {
	return config.DB
}
func (config *Config) GetServerConfig() *ServerConfig {
	return config.Server
}

func (config *Config) GetENVConfig() *ENVConfig {
	return config.ENV
}
