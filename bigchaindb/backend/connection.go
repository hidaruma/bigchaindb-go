package backend

type Connection struct {
	Host string
	Port int
	DBName string
	ConnectionTimeout int
	MaxTries int
	Kargs interface{}
}

func (conn *Connection) Conn() {
	
}

func (conn * Connection) Run() {
	
}

func (conn *Connection) Connect() {
	var attempt int
	attempt = 0
	for ()
}

func Connect(backend , host, port, name, maxTries, connectionTimeout, replicaset, ssl, login, password, caCert, certfile, keyfile, keyfilePassphrase, crlFile) {
	
}

