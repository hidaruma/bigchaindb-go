package pipelines

import (
	"github.com/hidaruma/bigchaindb-go/multipipes"
	"github.com/hidaruma/bigchaindb-go/bigchaindb"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
	"log"
)

type Vote struct {
	Bigchain *bigchaindb.Bigchain
	LastVotedID int
	Counters int
	BlocksValidityStatus []string
	DummyTx *bigchaindb.Transaction
}

func (v *Vote) init(blockDict map[string]Block) {
	if !v.Bigchain.HasPreviousVote(blockDict) {

	}
}