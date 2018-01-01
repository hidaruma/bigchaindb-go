package mongodb

import (
	"time"
	"log"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/backend/changefeed"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/backend/utils"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/mongodb/connection"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/backend/exceptions"
)

type MongoDBChangeFeed struct {
	
}

func (mcf *MongoDBChangeFeed) RunForever() {
	for _, element := range mcf.Prefeed {
		mcf.Outqueue.Put(element)
	}
	var table
	table = mcf.Table

	var dbname
	dbname = mcf.Connection.DBName

	var lastTs
	lastTs = mcf.Connection.Run(mcf.Connection.Query().//TODO)

	for _, record := range RunChangeFeed(mcf.Connection, table, lastTs) {
		var isInsert bool
		var isDelete bool
		var isUpdate bool

		switch record["op"] {
			case 'i':
				isInsert = true
			case 'd':
				isDelete = true
			case 'u':
				isUpdate = true
			default:
				isInsert = false
				isDelete = false
				isUpdate = false
		}

		if isInsert && (mcf.Operation && ChangeFeed.Insert) {
			record['o'].Pop("_id", "")
			mcf.Outqueue.Put(record['o'])
		} else if isDelete && (mcf.Operation && ChangeFeed.Delete) {
			mcf.Outqueue.Put(record['o'])
		} else if isUpdate && (mcf.Operation && ChangeFeed.Update) {
			var doc
			doc = mcf.Connection.Conn[dbname][table].findOne({"_id": record["o2"]["_id"]}, {"_id": false})
			mcf.Outqueue.Put(doc)
		}
		log.Debug()
	}
}

func GetChangeFeed(connection, table, operation, , prefeed) {

	
}

var FeedStop bool
FeedStop = false

func RunChangeFeed(conn, table, lastTs) {

	for {
		conn._conn = ""
		var namespace string
		namespace = conn.DBName + "." + table
		var query string
		query = conn.Query()
	
		var cursor
		cursor = conn.Run(query)
		log.Debug()

		for {
			if cursor.alive != true {
				break
			} else {
				var record 
				record = cursor.Next()
				err := yield record
				var lastTs
				lastTs = record["ts"]
				if err != nil {
					return
				}
			}
		}
		if err != nil {
			log.Println()
			time.Sleep(1)
		}
	}
}