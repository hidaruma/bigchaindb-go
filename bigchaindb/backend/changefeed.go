package backend

import (
	"log"
	"os"
	"runtime"
	"fmt"
	"strings"
	"github.com/hidaruma/bigchaindb-go/multipipes"
	"github.com/hidaruma/bigchaindb-go/bigchaindb"
)


type ChangeFeed struct {
	*multipipes.Node
	Table string
	Operation string
	Prefeed
	Connection Connection
}

func (cf *ChangeFeed) INSERT() int {
	return 1
}

func (cf *ChangeFeed) DELETE() int {
	return 2
}

func (cf *ChangeFeed) UPDATE() int {
	return 4
}

func (cf *ChangeFeed) init(table string, operation string, a interface{}, prefeed , connection Connection) {
	cf.Name = "changefeed"
	cf.Prefeed = prefeed
	cf.Table = table
	cf.Operation = operation
	if connection != nil {
		cf.Connection = connection
	} else {
		cf.Connection = Connect(bigchaindb.Config["database"])
	}

}

func (cf *ChangeFeed) RunForever() {
	
}

func (cf *ChangeFeed) RunChangeFeed() {
	
}

func GetChangeFeed(connection, table, operation, , prefeed) {
	
}