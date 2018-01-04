package bigchaindb

import (
	"os"
	"encoding/json"
	"fmt"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
	"github.com/hidaruma/bigchaindb-go/bigchaindb"
	"path/filepath"
	"log"
	"encoding/json"
	"reflect"
	"io/ioutil"
)

var BIGCHAINDBCONFIGPATH string = os.Getenv("BIGCHAINDB_CONFIG_PATH")
var CONFIGDEFAULTPATH string = filepath.Join(BIGCHAINDBCONFIGPATH, filepath.Join(os.ExpandEnv("$HOME"), ".bigchaindb"))

const (
	CONFIGPREFIX string = "BIGCHAINDB"
	CONFIGSEP string = "_"
)


type Mapping map[string]interface{}

func (m *Mapping) Get(k string, v interface{}) interface{} {
	dict := m.(map[string]interface{})
	if dict[k] != nil {
		return dict[k]
	} else {
		return v
	}
}

func MapLeafs(function func(a ...interface{} ) interface{} , mapping Mapping) Mapping {
	var path
	path = ""
	return inner(mapping, path, function)
}

func inner(mapping Mapping, path string, function func(a ...interface{}) interface{}) Mapping {
	for key, val := range mapping {
		if v, ok := val.(Mapping); ok == true {
			inner(v, path + key, function)
		} else {
			map[string]interface{}(mapping)[key] = function(val, path+key)
		}
	}
	return mapping
}

func Update(d Mapping, u Mapping) Mapping {
		for k, v := range u {
			if vVal, ok := v.(map[string]interface{}); ok {
				var defaultV interface{}
				defaultV = nil
				r := Update((d.Get(k, defaultV).(Mapping)), vVal)
				d[k] = r
			} else {
				d[k] = u[k]
			}
		}
	return d
}

func fileConfig(filename string) *Config {
	var config *Config
	if filename == "" {
		filename = CONFIGDEFAULTPATH
	}
	log.Println("fileConfig() will try to open" + filename + ".")
	filepathAbs, _ := filepath.Abs(filename)
	f, err := ioutil.ReadFile(filepathAbs)
	if err != nil {
		log.Fatal("File error: %v\n", err)
	}
	err = json.Unmarshal(f, config)
	if err != nil {
		log.Fatal(common.ConfigurationError())
	}
	log.Println("Configuration loaded from" +filename + ".")

	return config
}

func EnvConfig(config *Config) Mapping {
	return MapLeafs(loadFromEnv(value, path), config)
}

func loadFromEnv(value interface{}, path string)  []string {
	var varName string
	varName = filepath.Join(CONFIGPREFIX + )
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

func IsConfigured(config Config) bool {
	if config.Get("CONFIGURED", false) {
		return true
	} else {
		return false
	}
}

func Autoconfigure(filename string, config Config, force bool) *Config {
	if !force && IsConfigured() {
		log.Println("System already configured. skipping autoconfiguration.")
		return nil
	}
	var newconfig *Config
	newconfig = Update(Mapping(newconfig), Mapping(fileConfig(filename)).(*Config)
	newconfig = EnvConfig(newconfig)

	if config != nil {
		newconfig = Update(*newconfig.(Mapping), Mapping(config)).(*Config)
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

	if len(names) < 1 {
		return plugins
	}
	for _, name := range names {
		for _, entryPoint := range {
			plugins = append(plugins, entryPoint.Load())
		}
	}
	return plugins
}