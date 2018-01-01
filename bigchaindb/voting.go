package bigchaindb

import (
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
	"log"
)

const (
	VALID string = "valid"
	INVALID string = "invalid"
	UNDECIDED string = "undecided"
)

type Vote map[string]interface{}

type Voting struct {
	
}

func (v *Voting) BlockElection(block Block, votes []Vote, keyring string) map[string]interface{} {

	var eligibleVoters []string
	eligibleVoters = block.Voters
	var nVoters int
	nVoters = len(eligibleVoters)
	var eligibleVotes []Vote
	var ineligibleVotes []Vote
	eligibleVotes, ineligibleVotes = v.PartitionEligibleVotes(votes, eligibleVoters)
	var byVoter map[string]Vote
	byVoter = v.DedupeByVoter(eligibleVotes)
	var results map[string]interface{}
	results["block_id"] = block.ID
	results["status"] = v.DecideVotes(nVoters, results["counts"])
	results["ineligible"] = ineligibleVotes
	return results	
}

func (v *Voting) PartitionEligibleVotes(votes []Vote, eligibleVoters []string) ([]Vote, []Vote) {
	var eligible []Vote
	var ineligible []Vote

	for _, vote := range votes {
		if stringInSlice(vote.NodePubkey, eligibleVoters) {
			if v.VerifyVoteSignature(vote) {
					eligible = append(eligible, vote)
					continue
			}
		} else {
			ineligible = append(ineligible, vote)
		}
	}
	return eligible, ineligible
}

func (v *Voting) DedupeByVoter(eligibleVotes []Vote) map[string]Vote{
	var byVoter map[string]Vote
	for _, vote := range eligibleVotes {
		var pubkey string
		pubkey = vote.NodePubkey
		if StringInSlice(pubkey, byVoter) {
			log.Fatal(CriticalDuplicateVote(pubkey))
		}
		byVoter[pubkey] = vote
	}
	return byVoter
}

func (v *Voting) CountVotes(cls, byVoter map[string]Vote) {
	var prevBlocks map[string]Vote
	var malformed []Vote

	for _, vote := range byVoter {
		if !v.VerifyVoteSchema(vote) {
			malformed = append(malformed, vote)
			continue
		}
		if vote.IsBlockValid() == true {
			prevBlocks[vote["previousBlock"]]++
		}
	}
	var nValid
	nValid = 0
	var prevBlock Block

	if len(prevBlocks) != 0 {
		prevBlock, nValid = prevBlocks.most_common()[0]
		delete(prevBlocks[prevBlock])
	}

	return {
		""
	}
}

func (v *Voting) DecideVotes(cls, nVoters, nValid, nInvalid) {
	
}

func (v *Voting) VerifyVoteSignature(cls, vote) {
	
}

func (v *Voting) VerifyVoteSchema(cls, vote) bool {
	
}
