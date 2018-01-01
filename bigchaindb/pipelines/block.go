package pipelines

import (
	"github.com/hidaruma/bigchaindb-go/bigchaindb/backend"
	"github.com/hidaruma/bigchaindb-go/multipipes"
	"github.com/hidaruma/bigchaindb-go/bigchaindb"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
	"log"
)

type Tx struct {
	*bigchaindb.Transaction
}

type BlockPipeline struct {
	Bigchain *bigchaindb.Bigchain
	Txs []Tx
}

func (bp *BlockPipeline) init() {
	bp.Bigchain = bigchaindb.Bigchain{}
	bp.Txs = txCollector()
}

func (bp *BlockPipeline) FilterTx(tx Tx) map[string]string {
	if tx["assignee"] == bp.Bigchain.Me {
		delete(tx, "assignee")
		delete(tx, "assignee_timestamp")
		return tx
	}
}

func (bp *BlockPipeline) ValidateTx(tx Tx) *bigchaindb.Transaction {
	tx, err := bigchaindb.Transaction.FromDict(tx)
	if err != nil {
		log.Println(common.ValidationError())
	}
	if !bp.Bigchain.IsNewTransaction(tx.ID) {
		bp.Bigchain.DeleteTransaction(tx.ID)
		return nil
	}
	if tx.Operation == bigchaindb.Transaction.GENESIS() {
		log.Println(common.GenesisBlockAlreadyExistsError())
	}
	err = tx.Validate(bp.Bigchain)
	if err != nil {
		log.Println("Invalid tx: %v", err)
		bp.Bigchain.DeleteTransaction(tx.ID)
		return nil
	}
	return tx
}

func (bp *BlockPipeline) Create(tx Tx, timeout bool) {
	var txs []Tx
	txs =
	if len(txs) == 1000 || (timeout && txs != nil) {
		var block bigchaindb.Block{}
		block = bp.Bigchain.CreateBlock(txs)
		txs = txCollector()
		return block
	}
}

func createPipline() *BlockPipeline {
	var blokPipeline *BlockPipeline


	return blokPipeline
}

func Start() *BlockPipeline {
	var pipeline *BlockPipeline
	pipeline = createPipeline()
	pipeline.setup(GetChangefeed())
	pipeline.start()
	return pipeline
}