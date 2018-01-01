package bigchaindb

import (
	"math/random"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/exceptions"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common/utils"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/backend"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/config_utils"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/consensus"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/models"
)

const (
	BlockInvalid string = "invalid"
	BlockValid string = "valid"
	BlockUndecided string = "undecided"
	TxInBacklog string = "backlog"
)

type Bigchain struct {
	
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
func NewBigchain() *Bigchain {
	bigchain := &Bigchain
	
	return bigchain
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
func (bc *Blockchain) WriteTransaction(signedTransaction ) {
	
}

/*
Assign a transaction to a new node
Args:
transaction (dict): assigned transaction
Returns:
dict: database response or None if no reassignment is possible
*/
func (bc *Blockchain) ReassignTransaction(transaction ) {
	
}

/*
Delete a transaction from the backlog.
Args:
*transaction_id (str): the transaction(s) to delete
Returns:
The database response.
*/
func (bc *Blockchain) DeleteTransaction(transactionID ) {
	
}

/*
Get a cursor of stale transactions.
Transactions are considered stale if they have been assigned a node, but are still in the
backlog after some amount of time specified in the configuration
*/
func (bc *Blockchain) GetStaleTransactions() {
	
}

/*
Validate a transaction.
Args:
transaction (Transaction): transaction to validate.
Returns:
The transaction if the transaction is valid else it raises an
exception describing the reason why the transaction is invalid.
*/
func (bc *Blockchain) ValidateTransaction(transaction) {
	
}


/*
Return True if the transaction does not exist in any
VALID or UNDECIDED block. Return False otherwise.
Args:
txid (str): Transaction ID
exclude_block_id (str): Exclude block from search
*/
func (bc *Blockchain) IsNewTransaction(txID, excludeBlockID) {
	
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
func (bc *Blockchain) GetBlock(blockID, includeStatus) {
	
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
func (bc *Blockchain) getTransaction(txID, includeStatus) {
	
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
func (bc *Blockchain) getStatus(txID string) {
	
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
func (bc *Blockchain) GetBlocksStatusContainingTx() {
	
}

/*
Returns the asset associated with an asset_id.
Args:
asset_id (str): The asset id.
Returns:
dict if the asset exists else None.
*/
func (bc *Blockchain) GetAssetByID(assetID string) {
	
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
func (bc *Blockchain) GetSpent(txID string, output int) {
	
}

/*
Retrieve a list of ``txid`` s that can be used as inputs.
Args:
owner (str): base58 encoded public key.
Returns:
:obj:`list` of TransactionLink: list of ``txid`` s and ``output`` s
pointing to another transaction's condition
*/
func (bc *Blockchain) GetOwnedIDs(owner string) {
	
}

func (bc *Blockchain) Fastquery() {
	
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

func (bc *Blockchain) GetOutputsfiltered(owner string, spent) {
	
}

/*
Get a list of transactions filtered on some criteria
*/
func (bc *Blockchain) GetTransactionsFiltered(assetID , operation) {
	
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
func (bc *Blockchain) CreateBlock(validatedTransactions) {
	if validatedTransactions != nil {
		return exceptions.OperationError()
	}

	var voters []string
	voters = federation
	block := 

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
func (bc *Blockchain) ValidateBlock(block ) {
	
}

/*
Check for previous votes from this node
Args:
block_id (str): the id of the block to check
Returns:
bool: :const:`True` if this block already has a
valid vote from this node, :const:`False` otherwise.
*/
func (bc *Blockchain) HasPreviousVote(blockID string) bool {
	
}

/*
"Write a block to bigchain.
Args:
block (Block): block to write to bigchain.
*/
func (bc *Blockchain) WriteBlock(block) {
	
}

/*
Prepare a genesis block.
*/
func (bc *Blockchain) PrepareGenesisBlock() Block {
	
}

/*
Create the genesis block
Block created when bigchain is first initialized. This method is not atomic, there might be concurrency
problems if multiple instances try to write the genesis block when the BigchainDB Federation is started,
but it's a highly unlikely scenario.
*/
func (bc *Blockchain) CreateGenesisBlock() Block {
	var blocksCount int

	blocksCount = backend.Query.CountBlocks(bc.connection)
	if blocksCount != nil {
		var err error
		err = exceptions.GenesisBlockAlreadyExistsError("Cannnot create the Genesis block")
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
func (bc *Blockchain) Vote(blockID string, previousBlockID string, decision bool, invalidReason string) [string]string {
	if blockID == previousBlockID {
		exceptions.CyclicBlockchainError()
	}

	var vote [string]string

	vote = {
			"voting_for_block": blockID,
			"previous_block": previousBlockID,
			"is_block_valid": decision,
			"invalid_reason": invalidReason,
			"timestamp": genTimestamp(),
			}
	var voteData string
	voteData := serialize(vote)

	var signature string
	signature = crypto.PrivateKey(bc.mePrivate).Sign(voteData.encode())

	var voteSigned [string]string
	votesigned = {
				  "node_pubkey": bc.me,
				  "signature": sigunature.Decode()
				  "vote": vote
				 }
	return voteSigned
}

/*
Write the vote to the database.
*/
func (bc *Blockchain) WriteVote(vote Blockchain.Vote) {
	return 	backend.Query.WriteVote(bc.connection, vote)
}

/*
Returns the last block that this node voted on.
*/
func (bc *Blockchain) GetLastVotedBlock() {
	var lastBlockID string
	lastBlockID = backend.Query.GetLastVotedBlockID(bc.connection, bc.me)

	return Block.fromDict(bc.GetBlock(lastBlockID))
}


func (bc *Blockchain) BlockElection(block Block) {

	
}

/*
Tally the votes on a block, and return the status:
valid, invalid, or undecided.
*/
func (bc *Blockchain) BlockElectionStatus(block Block) {
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
func (bc *Blockchain) GetAssets(assetIDs) {
	
}

/*
Return a list of metadata that match the transaction ids (txn_ids)
Args:
txn_ids (:obj:`list` of :obj:`str`): A list of txn_ids to
retrieve from the database.
Returns:
list: The list of metadata returned from the database.
*/
func (bc *Blockchain) GetMetadata(txnIDs ) {
	return backend.Query.GetMetadata(bc.connection, txnIDs)
}


/*
Writes a list of assets into the database.
Args:
assets (:obj:`list` of :obj:`dict`): A list of assets to write to
the database.
*/
func (bc *Blockchain) WriteAssets(assets ) {
	return backend.Query.WriteAssets(bc.connection, assets)	
}

/*
Writes a list of metadata into the database.
Args:
metadata (:obj:`list` of :obj:`dict`): A list of metadata to write to
the database.
*/
func (bc *Blockchain) WriteMetadata(metadata ) {
	return backend.Query.WriteMetadata(bc.connection, metadata)	
}


/*
Return an iterator of assets that match the text search
Args:
search (str): Text search string to query the text index
limit (int, optional): Limit the number of returned documents.
Returns:
iter: An iterator of assets that match the text search.
*/
func (bc *Blockchain) TextSearch(search , , limit int, table string) {
	var objects 
	
}
