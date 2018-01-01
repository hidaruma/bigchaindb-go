package backend

import (
	"log"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
	"github.com/hidaruma/bigchaindb-go/bigchaindb"
)

const (
	TABLES []string{"bigchain", "backlog", "votes", "assets", "metadata"}
	VALIDLANGUAGES []string{"danish", "dutch", "english", "finnish", "french", "german", "hungarian", "italian", "norwgian", "portuguese", "romanian", "russian", "spanish", "swedish", "turkish", "none", "da", "nl", "en", "fi", "fr", "de", "hu", "it", "nb", "pt", "ro", "ru", "es", "sv", "tr"}
)

type Schemaer interface{
	CreateDatabase(*Connection, string)
	CreateTables(*Connection, string)
	CreateIndexes(*Connection, string)
	DropDatabase(*Connection, string)
}
func InitDatabase(connection *Connection, dbname string) {
	var conn *Connection
	if connection != nil {
		conn = connection
	} else {
		conn = Connect()
	}
	var dbn string
	if dbname != "" {
		dbn = dbname
	} else {
		dbn = bigchaindb.Config["database"]["name"]
	}

	CreateDatabase(conn, dbn)
	CreateTables(conn, dbn)
	CreateIndexes(conn, dbn)
}

func ValidateLanguageKey(obj map[string]interface{}, key string) {
	var backend string
	backend = bigchaindb.Config["database"]["backend"]

	if backend == "mongodb" {
		var data interface{}
		data = obj[key]
		switch data.(type) {
		case map[string]string:
			ValidateAllValuesForKey(data, "language", ValidateLAnguage)
		}
	}
}

func ValidateLanguage(value string) bool {
	var errorStr = "MongoDB does not support text search for the " +
		"language " + value + ". If you do not understand this error " +
		"message then please rename key/field \"language\" to " +
		"something else like \"lang\"."

	for _, language := range VALIDLANGUAGES {
		if value == language {
			return true
		}
	}
	log.Println(ValidationError(errorStr))
	return false
}