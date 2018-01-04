package bigchaindb

type BaseConsensusRules struct {
	Voting *Voting
}

func (b *BaseConsensusRules) ValidateTransaction(bigchain Bigchain, transaction Transaction) *Transaction {
	return transaction.Validate(bigchain)
}

func (b *BaseConsensusRules) ValidateBlock(bigchain Bigchain, block Block) bool {
	return block.Validate(bigchain)
}