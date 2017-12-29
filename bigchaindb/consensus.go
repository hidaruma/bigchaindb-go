package consensus

import (
	"github.com/hidaruma/bigchaindb-go/bigchaindb/voting"
)

type BaseConsensusRules struct {
	voting voting.Voting
}

func (b *BaseConsensusRules) ValidateTransaction(bigchain, transaction) {
	
	return 
}

func (b *BaseConsensusRules) ValidateBlock(bigchain, block) {

	return 
}