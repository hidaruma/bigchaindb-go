package bigchaindb

import (
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/backend"
)


type Transaction struct {
	*common.Transaction
}

/*
Validate transaction spend
Args:
bigchain (Bigchain): an instantiated bigchaindb.Bigchain object.
Returns:
The transaction (Transaction) if the transaction is valid else it
raises an exception describing the reason why the transaction is
invalid.
Raises:
ValidationError: If the transaction is invalid
*/
func (t *Transaction) Validate(bigchain Bigchain) {
	var inputConditions 
	if t.Operation == Transaction.TRANSEFER() {
		var inputTxs
		for _, input := range t.inputs {
			var inputTxID
			inputTxID = input.Fulfills.TxID
			var inputTx
			var status
			var includeStatus
			includeStatus = true
			inputTx, status = bigchain.GetTransaction(inputTxID, includeStatus)

			if inputTx == nil {
				
			}
			if status != bigchain.TxValid {
				
			}
			var spent
			spent = bigchain.GetSpent(inputTxID, input.Fulfills.Output)
			if spent != nil && spent.ID != t.ID {
				
			}
			var output
			output = inputTx.Outputs[input.Fulfills.Output]
			inputConditions = append(inputConditions, output)
			inputTxs = append(inputTxs, inputTx)
		}
		var links
		for _, i := range inputs {
			links = append(links, i.Fulfills.ToURL())
		}
		if len(links) != len(set(links)) {
			
		}

		var assetID
		assetID = Transaction.GetAssetID(inputTxs)
		if assetID != t.Asset["ID"] {
			
		}
		var inputAmount
		var outputAmount
		for _, inputCondition := range inputConditions {
			
		}
		for _, outputCondition := range t.Outputs {
			
		}

		if outputAmount != inputAmount {
			
		}
		return t
	}
}

func (t *Transaction) FromDict(cls, txBody) {
	
}

func (t *Transaction) FromDB(cls, bigchain, txDict) {
	
}

/*
Bundle a list of Transactions in a Block. Nodes vote on its validity.
Attributes:
transaction (:obj:`list` of :class:`~.Transaction`):
Transactions to be included in the Block.
node_pubkey (str): The public key of the node creating the
Block.
timestamp (str): The Unix time a Block was created.
voters (:obj:`list` of :obj:`str`): A list of a federation
nodes' public keys supposed to vote on the Block.
signature (str): A cryptographic signature ensuring the
integrity and validity of the creator of a Block.
*/
type Block struct {
	ID int
	Transactions []Transaction
	NodePubkey string
	Timestamp string
	Voters []string
	Signature string	
}

func (b *Block) Eq(other) {
	
}

/*
Validate the Block.
Args:
bigchain (:class:`~bigchaindb.Bigchain`): An instantiated Bigchain
object.
Note:
The hash of the block (`id`) is validated on the `self.from_dict`
method. This is because the `from_dict` is the only method in
which we have the original json payload. The `id` provided by
this class is a mutable property that is generated on the fly.
Returns:
:class:`~.Block`: If valid, return a `Block` object. Else an
appropriate exception describing the reason of invalidity is
raised.
Raises:
ValidationError: If the block or any transaction in the block does
not validate
*/

func (b *Block) Validate(bigchain Bigchain) *Block {
	b.ValidateBlock(bigchain)
	b.ValidateBlockTransactions(bigchain)
	return b
}


/*
Validate the Block without validating the transactions.
Args:
bigchain (:class:`~bigchaindb.Bigchain`): An instantiated Bigchain
object.
Raises:
ValidationError: If there is a problem with the block
*/
func (b *Block) ValidateBlock(bigchain Bigchain) {
	if b.NodePublkey {
		
	}
	if !b.IsSignatureValid() {
		
	}

	var txIDs []string
	for _, tx := range b.Transactions {
		txIDs = append(txIDs, tx.ID)
	}
	if len(txIDs) != len(txIDs) {
		
	}
}

/*
Validate Block transactions.
Args:
bigchain (Bigchain): an instantiated bigchaindb.Bigchain object.
Raises:
ValidationError: If an invalid transaction is found
*/
func (b *Block) ValidateBlockTransactions(bigchain Bigchain) {
	for _, tx := range b.Transactions {
		bigchain.ValidateTransaction(tx)
	}
}

/*
Create a signature for the Block and overwrite `self.signature`.
Args:
private_key (str): A private key corresponding to
`self.node_pubkey`.
Returns:
:class:`~.Block`
*/
func (b *Block) Sign(privateKey string) *Block {
	var blockBody 
	blockBody = b.ToDict()
	var blockSerialized
	blockSerialized = Serialize(blockBody["block"])
	var privateKey string
	privateKey = PrivateKey(privateKey)
	b.Signature = privateKey.Sign(blockSerialized.Encode()).decode()
	return b
}

/*
Check the validity of a Block's signature.
Returns:
bool: Stating the validity of the Block's signature.
*/
func (b *Block) IsSignatureValid() bool {
	var block
	block = b.ToDict()["block"]

	var blockSerialized
	blockSerialized = serialize(block).Encode()
	var publicKey string
	publicKey = PublicKey(block["node_pubkey"])

	
}

