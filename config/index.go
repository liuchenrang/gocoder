package config

import (
	"io/ioutil"
	"github.com/op/go-logging"
	"gopkg.in/yaml.v2"
)

const APP_SECRET = "1231234"

var (
	log     = logging.MustGetLogger("controller")
	Project ServerConfig
)

type ServerConfig struct {
	Port           int    `yaml:"port"`
	MsgTail        string `yaml:"msg_tail"`
	Database       string `yaml:"database"`
	RabcModel       string `yaml:"rabcModel"`
	DatabaseType   string `yaml:"database_type"`
	DaoImport []string `yaml:"daoImport"`
	IDaoImport []string `yaml:"idaoImport"`
	IServiceImport []string `yaml:"iserviceImport"`
	ServiceImport []string `yaml:"serviceImport"`
}

func NewServerConfig(yamlFile string) ServerConfig {
	server := ServerConfig{}
	content, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, &server)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("config value %+v,\r\n", server)
	return server
}
