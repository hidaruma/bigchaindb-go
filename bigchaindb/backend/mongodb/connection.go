package mongodb

import (
	"time"
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/backend"
	"github.com/hidaruma/bigchaindb-go/bigchaindb"
	"regexp"
	"github.com/hyperledger/fabric/common/tools/cryptogen/ca"
)
//https://www.compose.com/articles/connect-to-mongo-3-2-on-compose-from-golang/
func ClientWithSSL(host string, port string, username string, password string, dbname string, connectionTimeout int, caCert string, certfile string, keyfile string, keyfilePassphrase string, crlfile string) *mgo.Session {
	roots := x509.NewCertPool()
	if ca, err := ioutil.ReadFile(caCert); err != nil {
		roots.AppendCertsFromPEM(ca)
	}
	tlsConfig := &tls.Config{}
	tls.Config.RootCAs = roots
	var addr string
	addr = host + ":" + port
	var addrs []string
	addrs = append(addrs, addr)
	var dialInfo = mgo.DialInfo{
		Username: username,
		Password: password,
		Addrs:    addrs,
		Timeout:  connectionTimeout * time.Millisecond,
		Database: dbname,
		Mechanism: "MONGODB-X509",
	}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	return session
}

func Client(host string, port string, username string, password string, dbname string, connectionTimeout int) *mgo.Session {

}

var (
	MONGOOPTS = map[string]int{"socketTimeoutMS": 20000}
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
	mc.KeyfilePassphrase = keyfilePassphrase
	mc.Crlfile = crlfile
}

func (mc *MongoDBConnection) DB() string {
	return mc.Conn[mc.DBName]
}

func (mc *MongoDBConnection) Query() *bigchaindb.Lazy {
	return *bigchaindb.Lazy
}

func (mc *MongoDBConnection) Collection(name string) {
	return mc.Query()[mc.DBName][name]
}

func (mc *MongoDBConnection) Run(query Query) {
		
	//TODO
}

func (mc *MongoDBConnection) connect() {
	initializeReplicaSet(
						mc.Host,
						mc.Port,
						mc.ConnectionTimeout,
						mc.DBName,
						mc.Ssl,
						mc.Login,
						mc.Password,
						mc.CaCert,
						mc.Certfile,
						mc.Keyfile,
						mc.KeyfilePassphrase,
						mc.Crlfile,
						)
	if mc.CaCert == "" || mc.Certfile == "" || mc.Keyfile == "" || mc.Crlfile == "" {
		var client mgo.
		client = mg
	}

}

func initializeReplicaSet(host string, port string, connectionTimeout int, dbname string, ssl string, login string, password string, caCert string, certfile string, keyfile string, keyfilePassphrase string, crlfile string) {
	if (caCert == "") || (certfile == "") || (keyfile == "") {
		session, err := mgo.DialWithTimeout(host + ":" + port, connectionTimeout * time.Millisecond)
		if err != nil {
			panic(err)
		}
		defer session.Close()

		if login != "" && password != "" {
			credential := &mgo.Credential{Username:login, Password:password, Source:dbname}
			if err := session.Login(credential); err != nil {
				panic(err)
			}
			c := session.DB(dbname)
		}
	} else {
		log.Println("Authenticating to the database...")
		session, err := mgo.DialWithTimeout(host + ":" + port, connectionTimeout * time.Millisecond)
		if err != nil {
			panic(err)
		}
		defer session.Close()

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