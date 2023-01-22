package config

import (
	"fmt"
	"os"
	"strings"
)

var envVariables *EnvironmentVariables

type EnvironmentVariables struct {
	GoConfig         GoConfig
	DB               DBConfig
	YoutubeApiConfig YoutubeApiConfig
}

type GoConfig struct {
	BasePath   string
	ServerPort string
	Env        string
}

type YoutubeApiConfig struct {
	DeveloperKey   string
	Query          string
	MaxResults     int64
	PublishedAfter string
	Type           string
}

type DBConfig struct {
	Host            string `json:"host"`
	Port            uint   `json:"port"`
	User            string `json:"user"`
	Password        string `json:"password"`
	DBName          string `json:"DBName"`
	TableNamePrefix string `json:"tableNamePrefix"`
}

func GetEnv() EnvironmentVariables {
	if envVariables == nil {
		env := strings.ToLower(os.Getenv(EnvName))
		if env == "local" {
			fmt.Println("Loading default configs...")
			envVariables = readDefaultEnvVariables()
		}
	}
	return *envVariables
}

func readDefaultEnvVariables() *EnvironmentVariables {
	goConfigs := GoConfig{
		BasePath:   "fampay-assignment",
		ServerPort: "8081",
		Env:        "local",
	}
	dbVariables := DBConfig{
		Host:   "localhost",
		Port:   5432,
		User:   "brijesh",
		DBName: "assignment",
	}

	youtubeApiConfig := YoutubeApiConfig{
		DeveloperKey:   "AIzaSyCDRHwv0BQZddPkyl6UkyX6GvjHN5ttpUs",
		Query:          "Cricket",
		MaxResults:     10,
		PublishedAfter: "2023-01-21T03:32:16Z",
		Type:           "video",
	}

	return &EnvironmentVariables{
		GoConfig:         goConfigs,
		DB:               dbVariables,
		YoutubeApiConfig: youtubeApiConfig,
	}
}

const (
	EnvName = "ENV"
)
