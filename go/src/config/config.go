package config

import (
    "fmt"
    "github.com/xeipuuv/gojsonschema"
)

type Config struct {

	DB *DBConfig
	ENV  *ENVConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Charset  string
}

type ENVConfig struct {
	Language  string
	Library string
}

func NewConfig (configPath string) *Config {
	schemaLoader := gojsonschema.NewStringLoader("{}")
    documentLoader := gojsonschema.NewReferenceLoader(configPath)

    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        panic(err.Error())
    }
	if result.Valid() {
        fmt.Printf("The document is valid\n")
	} else {
        fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
            fmt.Printf("- %s\n", desc)
        }
	}
	return &Config{
		ENV: &ENVConfig{
			Language:  "python",
			Library:   "OpenCV",
		},
	}
}

func GetDBConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "database-service",
			Port:     3306,
			Username: "user",
			Password: "green",
			Charset:  "utf8",
		},
	}
}

//To be configured via config file i.e. config.json as a configmap
func GetENVConfig() *Config {
	return &Config{
		ENV: &ENVConfig{
			Language:  "python",
			Library:   "OpenCV",
		},
	}
}
