package bigchaindb

import (
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/backend"
	"time"
	"log"
	"encoding/json"
	"go/constant"
	"github.com/onsi/ginkgo/integration/_fixtures/watch_fixtures/B"
)

const (
	BlockInvalid string = "invalid"
	BlockValid string = "valid"
	BlockUndecided string = "undecided"
	TxInBacklog string = "backlog"
)
type Element map[string]interface{}

type Config map[string][string]interface{}

type Keyring []string

type Bigchain struct {
	Me string
	MePrivate string
	NodesExceptMe Keyring
	BacklogReassignDelay string
	Consensus BaseConsensusRules
	Connection *backend.Connection
	Statusd string
}

func (bc *Bigchain) BLOCKINVALID() string {
	return "invalid"
}

func (bc *Bigchain) BLOCKVALID() string {
	return "valid"
}

func (bc *Bigchain) BLOCKUNDECIDED() string {
	return "undecided"
}

func (bc *Bigchain) TXUNDECIDED() string {
	return "undecided"
}

/*
Initialize the Bigchain instance
A Bigchain instance has several configuration parameters (e.g. host).
If a parameter value is passed as an argument to the Bigchain
__init__ method, then that is the value it will have.
Otherwise, the parameter value will come from an environment variable.
If that environment variable isn't set, then the value
will come from the local configuration file. And if that variable
isn't in the local configuration file, then the parameter will have
its default value (defined in bigchaindb.__init__).
Args:
public_key (str): the base58 encoded public key for the ED25519 curve.
private_key (str): the base58 encoded private key for the ED25519 curve.
keyring (list[str]): list of base58 encoded public keys of the federation nodes.
connection (:class:`~bigchaindb.backend.connection.Connection`):
A connection to the database.
*/
func (bc *Bigchain) init(publicKey string, privateKey string, keyring []string, connection *backend.Connection, backlogReassignDelay string) {
	var config *Config
	config = Autoconfigure()
	if publicKey != nil {
		bc.Me = publicKey
	} else {
		bc.Me = config["keypair"]["public"]
	}
	if privateKey != nil {
		bc.MePrivate = privateKey
	} else {
		bc.MePrivate = config["keypair"]["private"]
	}
	if keyring != nil {
		bc.NodesExceptMe = keyring
	} else {
		bc.NodesExceptMe = config["keyring"]
	}
	if backlogReassignDelay == nil {
		backlogReassignDelay = config["backlog_reassign_delay"]
	}
	bc.BacklogReassignDelay = backlogReassignDelay

	var consensusPlugin Plugin
	consensusPlugin = config["consensus_plugin"]

	if consensusPlugin != nil {
		bc.Consensus = LoadConsensusPlugins(consensusPlugin)
	} else {
		bc.Consensus = BaseConsensusRules{}
	}

	if connection != nil {
		bc.Connection = connection
	} else {
		bc.Connection = backend.Connect(config["database"])
	}
}
func (bc *Bigchain) Federation() []string {
	var keys []string
	keys = append(bc.NodesExceptMe, bc.Me)
	return keys
}
/*
Write the transaction to bigchain.
When first writing a transaction to the bigchain the transaction will be kept in a backlog until
it has been validated by the nodes of the federation.
Args:
signed_transaction (Transaction): transaction with the `signature` included.
Returns:
dict: database response
*/
func (bc *Bigchain) WriteTransaction(signedTransaction *Transaction) {
	signedTransaction = signedTransaction.ToDict()
	var assignee string
	if bc.NodesExceptMe != nil {
		assignee = RandomChoiceString(bc.NodesExceptMe)
	} else {
		assignee = bc.Me
	}
	return backend.Query.WriteTransaction(bc.Connection, signedTransaction)
}

/*
Assign a transaction to a new node
Args:
transaction (dict): assigned transaction
Returns:
dict: database response or None if no reassignment is possible
*/
func (bc *Bigchain) ReassignTransaction(transaction *Transaction) {
	var otherNodes []*Transaction
	var newAssignee string
	otherNodes = bc.Federation.Difference([transaction["assignee"]])
	if otherNodes != nil {
		newAssignee = RandomChoiceString(otherNodes)
	} else {
		newAssignee = bc.Me
	}
	return backend.Query.UpdateTransaction(bc.Connection, transaction["id"],
		{"assignee": newAssignee, "assignment_timestamp": time.Now()})
}

/*
Delete a transaction from the backlog.
Args:
*transaction_id (str): the transaction(s) to delete
Returns:
The database response.
*/
func (bc *Bigchain) DeleteTransaction(transactionID string) {
	return backend.Query.DeleteTransaction(bc.Connection, transactionID)
}

/*
Get a cursor of stale transactions.
Transactions are considered stale if they have been assigned a node, but are still in the
backlog after some amount of time specified in the configuration
*/
func (bc *Bigchain) GetStaleTransactions() {
	return backend.Query.GetStaleTransaction(bc.Connection, bc.BacklogReassignDelay)
}

/*
Validate a transaction.
Args:
transaction (Transaction): transaction to validate.
Returns:
The transaction if the transaction is valid else it raises an
exception describing the reason why the transaction is invalid.
*/
func (bc *Bigchain) ValidateTransaction(transaction *Transaction) bool {
	return bc.Consensus.ValidateTransaction(transaction)
}


/*
Return True if the transaction does not exist in any
VALID or UNDECIDED block. Return False otherwise.
Args:
txid (str): Transaction ID
exclude_block_id (str): Exclude block from search
*/
func (bc *Bigchain) IsNewTransaction(txID string, excludeBlockID string)  bool {
	var blockStatuses []string
	blockStatuses = bc.GetBlocksStatusContainingTx(txID)
	delete(blockStatuses, excludeBlockID)
	for _, status := range blockStatuses {
		if status != bc.BLOCKINVALID() {
			return false
		}
	}
	return true
}

/*
Get the block with the specified `block_id` (and optionally its status)
Returns the block corresponding to `block_id` or None if no match is
found.
Args:
block_id (str): transaction id of the transaction to get
include_status (bool): also return the status of the block
the return value is then a tuple: (block, status)
*/
func (bc *Bigchain) GetBlock(blockID string, includeStatus bool) (map[string]Block, string){
	var blockDict map[string]Block
	blockDict = backend.Query.GetBlock(bc.Connection, blockID)
	if blockDict != nil {
		var assetIDs []string
		assetIDs = Block.GetAssetIDs(blockDict)
		var txnIDs []string
		txnIDs = Block.GetTxnIDs(blockDict)
		var assets []Asset
		assets = bc.GetAssets(assetIDs)
		var metadata Metadata
		metadata = bc.GetMetadata(txnIDs)
		blockDict = Block.CoupleAssets(blockDict, assets)
		blockDict = Block.CoupleMetadata(blockDict, metadata)

	}
	var status string
	if includeStatus {
		status = bc.BlockElectionStatus(blockDict)
		return blockDict, status
	} else {
		return blockDict, ""
	}
}


/*
Get the transaction with the specified `txid` (and optionally its status)
This query begins by looking in the bigchain table for all blocks containing
a transaction with the specified `txid`. If one of those blocks is valid, it
returns the matching transaction from that block. Else if some of those
blocks are undecided, it returns a matching transaction from one of them. If
the transaction was found in invalid blocks only, or in no blocks, then this
query looks for a matching transaction in the backlog table, and if it finds
one there, it returns that.
Args:
txid (str): transaction id of the transaction to get
include_status (bool): also return the status of the transaction
                      the return value is then a tuple: (tx, status)
Returns:
A :class:`~.models.Transaction` instance if the transaction
was found in a valid block, an undecided block, or the backlog table,
otherwise ``None``.
If :attr:`include_status` is ``True``, also returns the
transaction's status if the transaction was found.
*/
func (bc *Bigchain) GetTransaction(txID string, includeStatus bool) (*Transaction, string) {
	var response *Transaction
	var txStatus string
	var blocksValidityStatus map[string]string
	blocksValidityStatus = bc.GetBlocksStatusContainingTx(txID)
	var checkBacklog bool
	checkBacklog = trus

	if blocksValidityStatus != nil {
		for id, status := range blocksValidityStatus {
			if status != Bigchain.BLOCKINVALID() {
				blocksValidityStatus[id] = status
			}
		}
		if blocksValidityStatus != nil {
			checkBacklog = false

			for targetBlockID, _ := range blocksValidityStatus {
				if blocksValidityStatus[targetBlockID] == Bigchain.BLOCKVALID() {
					txStatus = bc.TXVALID
					break
				}

			}
		}
	}
	if checkBacklog {
		response = backend.Query.GetTransactionFromBacklog(bc.Connection, txID)
		if response != nil {
			txStatus = bc.TXINBACKLOG()
		}
	}
	if response != nil {
		if txStatus == bc.TXINBACKLOG() {
			response = Transaction.FromDict(response)
		} else {
			response = Transaction.FromDB(bc, response)
		}
	}
	if includeStatus {
		return response, txStatus
	} else {
		return response, nil
	}
}

/*
Retrieve the status of a transaction with `txid` from bigchain.
Args:
txid (str): transaction id of the transaction to query
Returns:
(string): transaction status ('valid', 'undecided',
or 'backlog'). If no transaction with that `txid` was found it
returns `None`
*/
func (bc *Bigchain) getStatus(txID string) string {
	var status string
	var includeStatus bool
	includeStatus = true
	_, status = bc.GetTransaction(txID, includeStatus)
	return status
}

/*
Retrieve block ids and statuses related to a transaction
Transactions may occur in multiple blocks, but no more than one valid block.
Args:
txid (str): transaction id of the transaction to query
Returns:
A dict of blocks containing the transaction,
e.g. {block_id_1: 'valid', block_id_2: 'invalid' ...}, or None
*/
func (bc *Bigchain) GetBlocksStatusContainingTx(txID string) map[int]string {
	var blocks []*Block
	blocks = backend.Query.GetBlocksStatusFromTransaction(bc.Connection, txID)

	if len(blocks) > 0 {
		var blocksValidityStatus map[int]string
		for _, block := range blocks {
			blocksValidityStatus[block.ID] = bc.BlockElectionStatus(block)
		}
		var validBlocksConter int
		for _, validity := range blocksValidityStatus {
			if validity == Bigchain.BLOCKVALID() {
				validBlocksConter++
			}
		}
		var blockIDs []string
		if validBlocksConter > 1 {
			for blockID, _ := range blocksValidityStatus {
				if blocksValidityStatus[blockID] == Bigchain.BLOCKVALID() {
					blockIDs = append(blockIDs, blockID)
				}
			}
			log.Println("%v, %v", txID, blockIDs)
		}
		return blocksValidityStatus
	} else {
		return nil
	}
}

/*
Returns the asset associated with an asset_id.
Args:
asset_id (str): The asset id.
Returns:
dict if the asset exists else None.
*/
func (bc *Bigchain) GetAssetByID(assetID string) *common.Asset {
	var cursor map[int]*common.Asset
	cursor = backend.Query.GetAssetByID(bc.Connection, assetID)
	if cursor != nil {
		return cursor[0]["asset"]
	}
	return nil
}


/*
Check if a `txid` was already used as an input.
A transaction can be used as an input for another transaction. Bigchain
needs to make sure that a given `(txid, output)` is only used once.
This method will check if the `(txid, output)` has already been
spent in a transaction that is in either the `VALID`, `UNDECIDED` or
`BACKLOG` state.
Args:
txid (str): The id of the transaction
output (num): the index of the output in the respective transaction
Returns:
The transaction (Transaction) that used the `(txid, output)` as an
input else `None`
Raises:
CriticalDoubleSpend: If the given `(txid, output)` was spent in
more than one valid transaction.
*/
func (bc *Bigchain) GetSpent(txID string, output int) *Transaction {
	var transactions []*Transaction
	transactions = backend.Query.GetSpent(bc.Connection, txID, output)
	var numValidTransactions int
	numValidTransactions = 0
	var nonInvalidTransactions []*Transaction
	for _, transaction := range transactions {
		var txn *Transaction
		var status string
		var includeStatus bool
		includeStatus = true
		txn, status = bc.GetTransaction(transaction["id"],includeStatus)
		if status == bc.TXVALID {
			numValidTransactions++
		}
		if numValidTransactions > 1 {
			log.Println(CriticalDoubleSpend(""))
		}
		if status != "" {
			transaction.Update({"metadata": txn.Metadata})
			nonInvalidTransactions = append(nonInvalidTransactions, transaction)
		}
	}
	if len(nonInvalidTransactions) > 0 {
		return Transaction.FromDict(nonInvalidTransactions[0])
	} else {
		return nil
	}
}

/*
Retrieve a list of ``txid`` s that can be used as inputs.
Args:
owner (str): base58 encoded public key.
Returns:
:obj:`list` of TransactionLink: list of ``txid`` s and ``output`` s
pointing to another transaction's condition
*/
func (bc *Bigchain) GetOwnedIDs(owner string) []*common.TransactionLink {
	var spent bool
	spent = false
	return bc.GetOutputsFiltered(owner, spent)
}

func (bc *Bigchain) Fastquery() {
	return Fastquery(bc.Connection, bc.Me)
}


/*
Get a list of output links filtered on some criteria
Args:
owner (str): base58 encoded public_key.
spent (bool): If ``True`` return only the spent outputs. If
``False`` return only unspent outputs. If spent is
not specified (``None``) return all outputs.
Returns:
:obj:`list` of TransactionLink: list of ``txid`` s and ``output`` s
pointing to another transaction's condition
*/

func (bc *Bigchain) GetOutputsFiltered(owner string, spent bool) []*common.TransactionLink {
	var outputs []*common.TransactionLink{}
	outputs = bc.Fastquery.GetOutputsByPublicKey(owner)
	switch spent {
	case nil:
		return outputs
	case true:
		return bc.Fastquery.FilterUnspentOutputs(outputs)
	case false:
		return bc.Fastquery.FilterSpentOutputs(outputs)
	}

}

/*
Get a list of transactions filtered on some criteria
*/
func (bc *Bigchain) GetTransactionsFiltered(assetID string, operation string) []*Transaction {
	var txIDs []string
	txIDs = backend.Query.GetTxIDsFiltered(bc.Connection, assetID, operation)
	var txs []*Transaction{}
	for _, txID := range txIDs {
		tx, status := bc.GetTransaction(txID, true)
		if status == bc.TXVALID() {
			txs = append(txs, tx)
		}
	}
	return txs
}

/*
Creates a block given a list of `validated_transactions`.
Note that this method does not validate the transactions. Transactions
should be validated before calling create_block.
Args:
validated_transactions (list(Transaction)): list of validated
                                          transactions.
Returns:
Block: created block.
*/
func (bc *Bigchain) CreateBlock(validatedTransactions []*Transaction) *Block {
	if len(validatedTransactions) > 0 {
		log.Println(OperationError())
	}
	var id int
	var voters []string
	voters = bc.Federation()
	block := *Block{
					id,
					validatedTransactions,
					bc.Me,
					GenTimestamp(),
					voters,
					}
	block.Sign(bc.MePrivate)
	return block
}

/*
Validate a block.
Args:
block (Block): block to validate.
Returns:
The block if the block is valid else it raises and exception
describing the reason why the block is invalid.
*/
func (bc *Bigchain) ValidateBlock(block *Block) bool {
	return bc.Consensus.ValidateBlock(bc, block)
}

/*
Check for previous votes from this node
Args:
block_id (str): the id of the block to check
Returns:
bool: :const:`True` if this block already has a
valid vote from this node, :const:`False` otherwise.
*/
func (bc *Bigchain) HasPreviousVote(blockID string) bool {
	var votes []Vote
	var el []Vote
	var keys []string
	keys = append(keys, bc.Me)
	votes = backend.Query.GetVotesByBlocksIDAndVoter(bc.Connection, blockID, bc.Me)
	el, _ = bc.Consensus.Voting.PartitionEligibleVotes(votes, keys)
	if len(el) > 0 {
		return true
	} else {
		return false
	}
}

/*
"Write a block to bigchain.
Args:
block (Block): block to write to bigchain.
*/
func (bc *Bigchain) WriteBlock(block *Block) {
	var assets []*common.Asset
	var blockDict map[string]*Block{}
	var metadatas []*common.Metadata
	assets, blockDict = block.DecoupleAssets()
	metadatas, blockDict = block.decoupleMetadata(blockDict)
	if len(assets) > 0 {
		bc.WriteAssets(assets)
	}
	if len(metadatas) > 0 {
		bc.WriteMetadata(metadatas)
	}
	return backend.Query.WriteBlock(bc.Connection, blockDict)
}

/*
Prepare a genesis block.
*/
func (bc *Bigchain) PrepareGenesisBlock() *Block {
	var blocksCount int
	blocksCount = backend.Query.CountBlocks(bc.Connection)

	if blocksCount > 0 {
		log.Println(common.GenesisBlockAlreadyExistsError("Cannot create the Genesis block"))
	}
	var block *Block
	block = bc.PrepareGenesisBlock()
	bc.WriteBlock(block)
}

/*
Create the genesis block
Block created when bigchain is first initialized. This method is not atomic, there might be concurrency
problems if multiple instances try to write the genesis block when the BigchainDB Federation is started,
but it's a highly unlikely scenario.
*/
func (bc *Bigchain) CreateGenesisBlock() *Block {
	var blocksCount int

	blocksCount = backend.Query.CountBlocks(bc.connection)
	if blocksCount != nil {
		log.Println(GenesisBlockAlreadyExistsError("Cannnot create the Genesis block"))
		return nil
	}
	block := bc.PrepareGenesisBlock()
	return block
}

/*
Create a signed vote for a block given the
attr:`previous_block_id` and the :attr:`decision` (valid/invalid).
Args:
block_id (str): The id of the block to vote on.
previous_block_id (str): The id of the previous block.
decision (bool): Whether the block is valid or invalid.
invalid_reason (Optional[str]): Reason the block is invalid
*/
func (bc *Bigchain) Vote(blockID string, previousBlockID string, decision bool, invalidReason string) *Vote {
	if blockID == previousBlockID {
		exceptions.CyclicBlockchainError()
	}

	var vote *Vote{
			VotingForBlock: blockID,
			PreviousBlockID: previousBlockID,
			IsBlockValid: decision,
			InvalidReason: invalidReason,
			Timestamp: GenTimestamp(),
			}
	var voteData string
	voteData := Serialize(vote)

	var signature string
	signature = crypto.PrivateKey(bc.MePrivate).Sign(voteData)

	var voteSigned *Vote{
				  NodePubkey: bc.Me,
				  Signature: sigunature.Decode()
				  Vote: vote
				 }
	return voteSigned
}

/*
Write the vote to the database.
*/
func (bc *Bigchain) WriteVote(vote *Vote) {
	return 	backend.Query.WriteVote(bc.Connection, vote)
}

/*
Returns the last block that this node voted on.
*/
func (bc *Bigchain) GetLastVotedBlock() {
	var lastBlockID string
	lastBlockID = backend.Query.GetLastVotedBlockID(bc.Connection, bc.Me)

	return Block.FromDict(bc.GetBlock(lastBlockID))
}


func (bc *Bigchain) BlockElection(block *Block) {
	switch block.(type) {
	case map[string]interface{}:
	default:
		block = block.ToDict()
	}
	var votes []*Vote
	votes = backend.Query.GetVotesByBlockID(bc.Connection, block.ID)
		return bc.Consensus.VotingBlockElection(block, votes, bc.Federation())
}

/*
Tally the votes on a block, and return the status:
valid, invalid, or undecided.
*/
func (bc *Bigchain) BlockElectionStatus(block *Block) {
	return bc.BlockElection(block)["status"]
}

/*
Return a list of assets that match the asset_ids
Args:
asset_ids (:obj:`list` of :obj:`str`): A list of asset_ids to
retrieve from the database.
Returns:
list: The list of assets returned from the database.
*/
func (bc *Bigchain) GetAssets(assetIDs []string) []*common.Asset{
	return backend.Query.GetAssets(bc.Connection, assetIDs)
}

/*
Return a list of metadata that match the transaction ids (txn_ids)
Args:
txn_ids (:obj:`list` of :obj:`str`): A list of txn_ids to
retrieve from the database.
Returns:
list: The list of metadata returned from the database.
*/
func (bc *Bigchain) GetMetadata(txnIDs []string) []*common.Metadata {
	return backend.Query.GetMetadata(bc.Connection, txnIDs)
}


/*
Writes a list of assets into the database.
Args:
assets (:obj:`list` of :obj:`dict`): A list of assets to write to
the database.
*/
func (bc *Bigchain) WriteAssets(assets []*common.Asset) {
	return backend.Query.WriteAssets(bc.Connection, assets)
}

/*
Writes a list of metadata into the database.
Args:
metadata (:obj:`list` of :obj:`dict`): A list of metadata to write to
the database.
*/
func (bc *Bigchain) WriteMetadata(metadata *common.Metadata) {
	return backend.Query.WriteMetadata(bc.Connection, metadata)
}


/*
Return an iterator of assets that match the text search
Args:
search (str): Text search string to query the text index
limit (int, optional): Limit the number of returned documents.
Returns:
iter: An iterator of assets that match the text search.
*/
func (bc *Bigchain) TextSearch(search , , limit int, table string) []interface{} {
	var objects []interface{}
	objects = backend.Query.TextSearch(bc.Connection, search, limit, table)
	var validObjects []interface{}
	for _, obj := range objects {
		tx, status := range bc.GetTransaction(obj["id"], true)
		if status == bc.TXVALID() {
			validObjects = append(validObjects, obj)
		}
	}
	return validObjects
}
