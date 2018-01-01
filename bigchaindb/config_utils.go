package bigchaindb

import (
	"os"
	"encoding/json"
	"fmt"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
	"github.com/hidaruma/bigchaindb-go/bigchaindb"
	"path/filepath"
	"log"
	"golang.org/x/net/html/atom"
)

const (
	CONFIGDEFAULTPATH string = filepath.Join("", ".bigchaindb")
	CONFIGPREFIX string = "BIGCHAINDB"
	CONFIGSEP string = "_"
)


type Mapping map[string]interface{}

type Config Mapping

func MapLeafs(function func(a ...interface{} ) interface{} , mapping Mapping) {
	return inner(mapping, )
}

func inner(mapping Mapping, path string, function func(a ...interface{}) interface{}) Mapping {
	for key, val := range mapping {
		switch val {
		case Mapping:
			inner(val, path + key, function)
		default:
			mapping[key] = function(val, path+key)
		}
	}
	return mapping
}

func Update(d Mapping,u Mapping) Mapping {
	for k, v := range u {
		switch v.(type) {
		case Mapping:
			r := Update(d[k], v)
		default:
			d[k] = u[k]
		}
	}
	return d
}

func fileConfig(filename string) Config {
	var config Config
	if filename == "" {
		filename = CONFIGDEFAULTPATH
	}
	log.Println("fileConfig() will try to open" + filename + ".")
	f, err := os.Open(filepath.Abs(filename))
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(f, config)
	if err != nil {
		log.Println(common.ConfigurationError(err))
	}
	log.Println("Configuration loaded from" +filename + ".")

	return config
}

func EnvConfig(config Config) Mapping {
	return MapLeafs(loadFromEnv(value, path), config)
}

func loadFromEnv(value interface{}, path string)  {

}

func UpdateTypes(config , reference, listSep ) {
	
}

func SetConfig(config *Config) {

	_ = Update(bigchaindb.Config, updateTypes(config, bigchaindb.Config))

}

func UpdateConfig(config ) {
	
}

func WriteConfig(config, filename string) {
	
}

func IsConfigured() bool {
	
}

func Autoconfigure(filename string, config *Config, force bool) {
	if !force && IsConfigured() {
		log.Println("System already configured. skipping autoconfiguration.")
		return
	}
	var newconfig Config
	newconfig = Config(Update(Mapping(newconfig), Mapping(fileConfig(filename))))
	newconfig = EnvConfig(newconfig)

	if config != nil {
		newconfig = Config(Update(Mapping(newconfig), Mapping(config)))
	}
	SetConfig(newconfig)
}

func LoadConsensusPlugin(name string) BaseConsensusRules {
	if name == "" {
		return BaseConsensusRules{}
	}

	var plugin Plugin
	for _, entryPoint := range {
		plugin = entryPoint
	}
}

func LoadEventsPlugins(names []string) []Plugin {
	var plugins []Plugin

	if names == "" {
		return plugins
	}
	for _, name := range names {
		for _, entryPoint := range {
			plugins = append(plugins, entryPoint.Load())
		}
	}
	return plugins
}