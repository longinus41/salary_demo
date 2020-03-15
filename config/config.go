package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

//MainConfig represents basic config
type MainConfig struct {
	Prime      int64 `json:"prime"`
	Generator  int64 `json:"generator"`
	PrivateKey int64 `json:"private_key"`
}

//Configs represents other configs
type Configs map[string]json.RawMessage

var configPath string = "../config/configs.json"

var conf *MainConfig
var confs Configs

var instanceOnce sync.Once

//LoadConfig loads configs from JSON
func LoadConfig(path string) (Configs, *MainConfig) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	mainConfig := &MainConfig{}
	err = json.Unmarshal(buf, mainConfig)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}
	allConfigs := make(Configs, 0)
	err = json.Unmarshal(buf, &allConfigs)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}

	return allConfigs, mainConfig
}

//Init is the init function
func Init() (*MainConfig, Configs) {
	if conf != nil {
		log.Printf("the config is already initialized")
	}
	instanceOnce.Do(func() {
		allConfigs, mainConfig := LoadConfig(configPath)
		conf = mainConfig
		confs = allConfigs
	})
	return conf, confs
}
