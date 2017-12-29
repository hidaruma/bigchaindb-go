package bigchaindb

import (
	"os"
	"encoding/json"
	"fmt"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
	"github.com/hidaruma/bigchaindb-go/bigchaindb"
	"path/filepath"
)

const (
	ConfigDefaultPath string = filepath.Join("", ".bigchaindb")
	ConfigPrefix string = "BIGCHAINDB"
	ConfigSep string = "_"
)

func MapLeafs(function, mapping) {
	
}

func Update(d ,u ) {
	
}

func fileConfig(filename string) {
	
}

func EnvConfig() {}

func UpdateTypes(config , reference, listSep ) {
	
}

func SetConfig(config ) {
	
}

func UpdateConfig(config ) {
	
}

func WriteConfig(config, filename string) {
	
}

func IsConfigured() bool {
	
}

func Autoconfigure(finlename string, config , force) {
	
}

func LoadConsensusPlugin(name string) {
	
}

func LoadEventsPlugins(names []string) []string {
	var plugins []string

	if names != nil {
		return plugins
	}
	for _, entryPoint := range {
		plugins = append(plugins, )
	}
	return plugins
}