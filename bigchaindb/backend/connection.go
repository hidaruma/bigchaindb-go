package backend

import (
	"github.com/hidaruma/bigchaindb-go/bigchaindb"
	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"log"
)

const BACKENDS map[string]string{
	"mongodb":"bigchaindb.backend.mongodb.connection.MongoDBConnection",
"rethinkdb": "",
"boltdb": "",
}


type Connection struct {
	Host string
	Port int
	DBName string
	Ssl string
	Login string
	Password string
	CaCert string
	Certfile string
	Keyfile string
	KeyfilePassphrase string
	Crlfile string
	Replicaset string
	ConnectionTimeout int
	MaxTries int
	MaxTriesCounter []int
	Kwargs interface{}
	conn *Connection
	Connectioner
}

func (conn *Connection) init(host string, port int, dbname string, connectionTimeout int, maxTries int, kwargs interface{}) {
	var dbconf string
	dbconf = bigchaindb.Config["database"]

	if host != nil {
		conn.Host = host
	} else {
		conn.Host = dbconf["host"]
	}
	if port != nil {
		conn.Port = port
	} else {
		conn.Port = dbconf["port"]
	}
	if dbname != nil {
		conn.DBName = dbname
	} else {
		conn.DBName = dbconf["name"]
	}
	if connectionTimeout != nil {
		conn.ConnectionTimeout = connectionTimeout
	} else {
		conn.ConnectionTimeout = dbconf["connection_timeout"]
	}
	if maxTries != nil {
		conn.MaxTries = maxTries
	} else {
		conn.MaxTries = dbconf["max_tries"]
	}
	if maxTries != 0 {
		conn.MaxTriesCounter = Iterator(conn.MaxTries)
	} else {
		conn.MaxTriesCounter = repeat(0)
	}
	conn.conn = nil

}

func (conn *Connection) Conn() Connection.conn {
	if conn.conn == nil {
		conn.Connect()
	}
	return conn.conn
}

type Connectioner interface {
	Run(Query)
}

func (conn *Connection) Connect() {
	var attempt int
	attempt = 0
	for i := range conn.MaxTriesCounter {
		attempt++
		conn.conn, err := conn.connect()
		if err != nil {
			log.Println(err)
			if attmpt == conn.MaxTries {
				log.Fatal("Cannnot connect to the Database. Giving up.")
			}
		} else {
			break
		}
	}
}

func Connect(backend string, host string, port string , name string, maxTries int, connectionTimeout int, replicaset string, ssl string, login string, password string, caCert string, certfile string, keyfile string, keyfilePassphrase string, crlfile string) *Connection {
	if backend == nil {
		backend = bigchaindb.Config["database"]["replicaset"]
	}
	if host == nil {
		host = bigchaindb.Config["database"]["host"]
	}
	if port == nil {

	}
	var dbname string
	if name == nil {

	}
	if replicaset == nil {

	}
	if ssl == nil {

	}
	if login == nil {

	}
	if password == nil {

	}
	if caCert == nil {

	}
	if certfile == nil {

	}
	if keyfile == nil {

	}
	if keyfilePassphrase == nil {

	}
	if crlfile == nil {

	}

	var moduleName string
	var className string
	moduleName, _, className = BACKENDS

	//TODO
	var class *Connection{
		Host: host,
		Port: port,
		DBName: dbname,
		MaxTries: maxTries,
		ConnectionTimeout: connectionTimeout,
		Replicaset: replicaset,
		Ssl: ssl,
		Login: login,
		Password: password,
		CaCert: caCert,
		Certfile: certfile,
		Keyfile: keyfile,
		KeyfilePassphrase: keyfilePassphrase,
		Crlfile: crlfile,
	}
	log.Println("Connection: %v", class)

	return class
}
