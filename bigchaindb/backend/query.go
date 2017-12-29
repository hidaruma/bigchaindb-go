package backend

func WriteTransaction(connection , signedTransaction ) error {


	return 
}

func UpdateTransaction(connection, ) error {

	return 
}

func DeleteTransaction(connection, transactionID ) error {
	
}


func GetStateTransactions(connection, transactionID ) error {
	
}

func GetTransactionFromBlock(connection, transactionID , blockID ) error {
	
}

func GetTransactionFromBacklog(connection, transactionID ) error {
	
}

func GetBlocksStatusFromTransaction(connection, transactionID ) error {
	
}

func GetAssetByID(connection , assetID ) error {
	
}

func GetSpent(connection, transactionID, conditionID) error {
	
}

func GetSpendingTransactions(connection, inputs ) error {
	
}

func GetOwnedIDs(connection, owner) error {
	
}

func GetVotesByBlockIDAndVoter(connection, blockID, nodePubkey) error {
	
}

func GetVotesForBlocksByVoter(connection, BlockIDs, pubkey) error {
	
}

func WriteBlock(connection, block) error {
	
}

func GetBlock(connection, blockID) error {
	
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