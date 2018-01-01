package bigchaindb

import (
	"github.com/hidaruma/bigchaindb-go/bigchaindb/backend/query"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/utils"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common/transaction"
)

type FastQuery struct {
	
}

func NewFastQuery(connection, me) FastQuery {
	var fastquery FastQuery
	fastquery = {connection: connection, me: me}
	return fastquery
}

func (fq *FastQuery) FilterValidBlockIDs(blockIDs, includeUndecided bool) {
	
}