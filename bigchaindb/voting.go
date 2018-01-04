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

func (v *Voting) BlockElection(block *Block, votes []*Vote, keyring string) *Vote {

	var eligibleVoters []*Vote
	eligibleVoters = block.Voters
	var nVoters int
	nVoters = len(eligibleVoters)
	var eligibleVotes []Vote
	var ineligibleVotes []Vote
	eligibleVotes, ineligibleVotes = v.PartitionEligibleVotes(votes, eligibleVoters)
	var byVoter map[string]Vote
	byVoter = v.DedupeByVoter(eligibleVotes)

	var results *Vote
	results = v.CountVotes(byVoter)
	map[string]interface{}(*results)["block_id"] = block.ID
	map[string]interface{}(*results)["status"] = v.DecideVotes(nVoters, results["counts"])
	map[string]interface{}(*results)["ineligible"] = ineligibleVotes
	return results	
}

func (v *Voting) PartitionEligibleVotes(votes []*Vote, eligibleVoters []string) ([]Vote, []Vote) {
	var eligible []Vote
	var ineligible []Vote

	for _, vote := range votes {
		if StringInSlice(Mapping(vote)["node"], []interface{}(eligibleVoters)) {
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

func (v *Voting) CountVotes(byVoter map[string]*Vote) {
	var prevBlocks map[string]int
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

func (v *Voting) DecideVotes(nVoters int, nValid int, nInvalid int) string {
	if nInvalid *2 >= nVoters {
		return INVALID
	}
	if nValid *2 >= nVoters{
		return VALID
	}
	return UNDECIDED
}

func (v *Voting) VerifyVoteSignature(vote *Vote) bool {
	var signature string
	var pkBase58 string
	voteSignature := map[string]interface{}(*vote)["signature"]
	voteNodePubkey := map[string]interface{}(*vote)["node_pubkey"]
	if voteSignatureString, ok := voteSignature.(string); ok {
		signature = voteSignatureString
	} else {
		log.Fatal(ValueError())
	}
	if voteNodePubkeyString, ok := voteNodePubkey.(string); ok {
		pkBase58 = voteNodePubkeyString
	}
	var pubKey common.PublicKey
	pubKey = common.PublicKey(pkBase58)
	var body string
	voteVote := map[string]interface{}(*vote)["vote"]
	if voteVoteString, ok := voteVote.(map[string]string); ok {
		body = common.Serialize(voteVoteString)
	}

	return pubKey.Verify(body, signature)
}

func (v *Voting) VerifyVoteSchema(vote Vote) bool {
	VaridateVoteSchema(vote)

}
