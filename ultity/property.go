package ultity

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type DBConfig struct {
	DbServer   string
	Abc        string
	DbPort     string
	DbUser     string
	DbPassword string
	Database   string
}

type ServerConfig struct {
	Server       string
	PortBackend  string
	PortFrontend string
}

func ReadDBConfig() DBConfig {
	var configFile = "properties.ini"
	_, err := os.Stat(configFile)
	if err != nil {
		log.Fatal("property file is missing: ", configFile)
	}

	var config DBConfig
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

func ReadServerConfig() ServerConfig {
	var configFile = "properties.ini"
	_, err := os.Stat(configFile)
	if err != nil {
		log.Fatal("property file is missing: ", configFile)
	}

	var config ServerConfig
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
