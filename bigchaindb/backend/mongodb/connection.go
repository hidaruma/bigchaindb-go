package mongodb

import (
	"time"
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/backend"
	"regexp"
)

type MongoDBConnection struct {
	*backend.Connection
}

func (mc *MongoDBConnection) init(replicaset string, ssl string, login string, password string, caCert string, certfile string, keyfile string, keyfilePassphrase string, crlfile string, kwargs interface{}) {

	mc.Replicaset = replicaset
	mc.Ssl = ssl
	mc.Login = login
	mc.Password = password
	mc.CaCert = caCert
	mc.Certfile = certfile
	mc.Keyfile = keyfile
	mc.keyfilePassphrase = keyfilePassphrase
	mc.Crlfile = crlfile
}

func (mc *MongoDBConnection) DB() string {
	return mc.Conn[mc.DBName]
}

func (mc *MongoDBConnection) Query() {
	return Lazy()
}

func (mc *MongoDBConnection) Collection(name string) {
	return mc.Query()[mc.DBName][name]
}

func (mc *MongoDBConnection) Run(query Query) {
		
	//TODO
}

func InitializeReplicaSet(host string, port string, connectionTimeout int, dbname string, ssl string, login string, password string, caCert string, certfile string, keyfile string, keyfilePassphrase string, crlfile string) {
	if (caCert == "") || (certfile == "") || (keyfile == "") {
		var conn mgo. //todo

		if login != "" && password != "" {
			conn[dbname].authenticate(login, password)
		}
	} else {
		log.Println("Authenticating to the database...")
		conn[dbname].authenticate(login, "MON")
	}


}

func checkReplicaSet(conn *MongoDBConnection) {
	var options
	options = conn.Admin.Command("getCmdLineOpts")
	var replOpts []string
	replOpts = options["parsed"]["replication"]
	var replSetName string
	replSetName = replOpts.Get()
}

func waitForReplicaSetInitialization(conn *MongoDBConnection) {
	for {
		var logs = string
		logs = conn.Admin.Command("getLog", "rs")["log"]
		if regexp.Find("database writes are now permitted", logs) {
			return
		}
		time.Sleep(0.1)
	}
}