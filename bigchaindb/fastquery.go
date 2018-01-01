package bigchaindb

import (
	"github.com/hidaruma/bigchaindb-go/bigchaindb/backend"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
)

type FastQuery struct {
	*backend.Query
}

func (fq *FastQuery) init(connection backend.Connection, me string) {
	fq.Connection = connection
	fq.Me = me
}

func (fq *FastQuery) FilterValidBlockIDs(blockIDs []string, includeUndecided bool) {
	var votes []*Vote
	votes = backend.Query.GetVotesForBlocksByVoter(fq.Connection, blockIDs, fq.Me)
	for _, vote := range votes {
		}
	}
}

func (fq *FastQuery) FilterValidItems(items map[string]interface{}, blockIDKey func(b ) b[0]) map[string]interface{} {

}

func (fq *FastQuery) GetOutputsByPublicKey(publicKey string) []*common.TransactionLink {
	var res []*common.TransactionLink
	var txs []*common.Transaction
	res = backend.Query.GetOwnedIDs(fq.Connection, publicKey)
	for _, tx := range fq.FilterValidItems(res) {
		txs = append(txs, tx)
	}
	var txls []*common.TransactionLink
	for _, tx := range txs {
		for index, output := range tx["outputs"] {
			if ConditionDetailsHasOwner(output["condition"]["details"], publicKey) {
				txls = append(txls, *common.TransactionLink{tx.ID, index})
			}
		}
	}
	return txls
}

func (fq *FastQuery) FilterSpentOutputs(outputs []*common.TransactionLink) {
	var links map[string]string
	for _, output := range outputs {
		links[]
	}
	return
}

func (fq *FastQuery) FilterUnspentOutputs(outputs []*common.TransactionLink) {

	return
}