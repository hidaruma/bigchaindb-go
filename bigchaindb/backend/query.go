package backend

import (
	"github.com/hidaruma/bigchaindb-go/bigchaindb"
)

type Queryer interface {
	WriteTransaction(*Connection, *Transaction) error
	UpdateTransaction(*Connection, string) error
	DeleteTransaction(*Connection, string) error
	GetStateTransactions(*Connection, string) error
	GetTransactionFromBlock(*Connection, string, string) error
	GetTransactionFromBacklog(*Connection, string) error
	GetBlocksStatusFromTransaction(*Connection, string) error
	GetAssetByID(*Connection, string) error
	GetSpent(*Connection, string, string) error
	GetSpendingTransactions(*Connection, []*Input) error
	GetOwnedIDs(*Connection, string) error
	GetVotesByBlockIDAndVoter(*Connection, string, string) error
	GetVotesForBlockByVoter(*Connection, string, string) error
	WriteBlock(*Connection, *bigchaindb.Block) error
	GetBlock(*Connection, string) error
}

func WriteAssets(connection, assets) error {
	
}

func WriteMetadata(connection, metadata) error {
	
}

func GetAssets(connection, assetIDs) error {
	
}

func GetMetaData(connection, txnIDs) error {
	
}

func CountBlocks(connection) error {
	
}

func CountBacklog(connection) error {
	
}

func WriteVote(connection, vote) error {
	
}

func GetGenesisBlock(connection) error {
	
}

func GetLastVotedBlockID() error {
	
}

func GetTxIDsFiltered(connection, assetID, operation) error {
	
}

func GetNewBlocksFeed(connection, startBlockID) error {
	
}

func TextSearch(conn, search, , language , caseSensitive, diacriticSensitive, textScore, limit, table) error {

}