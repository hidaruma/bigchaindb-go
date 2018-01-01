package common

import (
	"time"
	"encoding/json"
	"strings"
	"regexp"
	"github.com/hidaruma/bigchaindb-go/bigchaindb"
	"log"
)
func GenTimestamp() string {
	return time.Now().Unix()
}

func Serialize(data [string]string) string {
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	return string(b)
}

func Deserialize(data string) interface{} {
	var ret interface{}
	err := json.Unmarshal(data, ret)
	if err != nil {
		log.Println(err)
	}
	return ret
}

func ValidateTxnObj(objName string, obj map[string]interface{}, key string, validationFun func()) {
	var backend
	backend = bigchaindb.config["database"]["backend"]

	if backend == "mongodb" {
		var data interface{}{}
		data = obj[key]
		switch data.(type) {
		case map[string]string:
				ValidateAllKeys(objName, data, validationFun)
				if err != nil {
					log.Println(err)
			}
		}
	}
}

func ValidateAllKeys(objName string, obj map[string]interface{}, validationFun func(interface{} ...) error) {
	for key, value := range obj {
		err := validationFun(objName, key)
		if err != nil {
			log.Println(err)
		}
		switch value.(type) {
			case map[string]interface{} :
				ValidateAllKeys(objName, value, validationFun)
		}
	}
}

func ValidateAllValuesForKey(obj map[string]interface{}, key string, validationFun func(interface{}) error) {
	for vkey, value := range obj {
		if vkey == key {
			err := validationFun(value)
			if err != nil {
				logPrintln(err)
			}
		} else {
			switch value.(type) {
			case map[string]interface{}:
				ValidateAllValuesForKey(value, key, validationFun)
			}
		}
	}
}

func ValidateKey(objName string, key string) error {
	if regexp.Find(`^[$|\.|\x00`, key) {
		var errorStr string
		errorStr = "Invalid key name " + key +" in "+  objName +" object. The ''key name cannot contain characters ''".", "$" or null characters"
		return exceptions.ValidationError(errorStr)
	} else {
		return nil
	}
}