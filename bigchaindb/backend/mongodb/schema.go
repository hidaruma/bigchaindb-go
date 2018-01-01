package mongodb

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/backend"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
)

type registerSchema backend.ModuleDispatchRegister(backend.Schema)

func CreateDatabase(connection *Connection, dbname string) {
	for _, dbexists := connection.Conn.DatabaseNames() {
		if dbname == dbexists {
			log.Fatal(backend.DatabaseAlreadyExists("Database " + dbname + " already exists."))
		}
	}
	log.Println("Create database " + dbname + ".")
	conn.Conn.GetDatabase(dbname)
}

func CreateTables(connection *Connection, dbname string) {
	for _, tableName := range backend.TABLES {
		log.Println("Create " + tableName + " table.")

		connection.Conn["dbname"].CreateCollection(tableName)
	}
}