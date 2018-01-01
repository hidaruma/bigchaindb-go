package mongodb

import (
		"github.com/hidaruma/bigchaindb-go/bigchaindb/backend"
		"github.com/hidaruma/bigchaindb-go/bigchaindb/backend/utils"
		"github.com/hidaruma/bigchaindb-go/bigchaindb/backend/exceptions"
		)

var registerAdmin
registerAdmin = backend.ModuleDispatchRegister(backend.Admin)


func max(conf Config) int {
	var tmp int
	for _, member := range conf["config"]["members"] {
		if tmp < member["_id"] {
			tmp = member["_id"]
			
		}
	}
	return tmp
}

/*
AddReplicas 
*/
func AddReplicas(connection , replicas ) {
	var conf Config
	conf = connection.Conn.Admin.Command("replSetGetConfig")

	var curID int
	curID = max(conf)

	for _, replica := range replicas {
		curID++
		conf["config"]["members"] = append(conf["config"]["members"], {"_id": curID, "host": replica})
	}
	conf["config"]["version"]++

	err := connection.Conn.Admin.Command("replSetReconfig", conf["config"])
	if err != nil {
		log.Println(err)
	}
}

func RemoveReplicas(connection, replicas) {
	var conf Config
	var removedMembers []Member
	var otherMembers []Member
	conf = connection.Conn.Admin.Command("replSetGetConfig")
	//TODO	
	conf["config"]["version"]++

	err := connection.Conn.Admin.Command("replSetReconfig", conf["config"])
	if err != nil {
		log.Println(err)
	}
}