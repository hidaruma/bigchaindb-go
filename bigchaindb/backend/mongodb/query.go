package mongodb

import (
	"time"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/backend"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
	"github.com/hyperledger/fabric/protos/peer"
	"log"
)

func WriteQuery(conn *MongoDBConnection, signedTransaction *common.Transaction) {
	connrun, err := conn.Run(conn.Collection("backlog").InsertOne(signedTransaction))
	if err != nil {
		log.Println(err)
		return
	}
	return connrun
}

func UpdateTransaction() {

}




func removeTextScore(asset *common.Asset) *common.Asset {
	delete(asset, "score")
	return asset
}