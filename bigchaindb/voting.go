package bigchaindb

type Voting struct {
	
}

func (v *Voting) BlockElection(cls , block , votes , keyring ) []string {
	var results []string

	return results	
}

func (v *Voting) PartitionElicibleVotes(cls, votes, eligibleVoters) ([]Vote, []Vote) {
	var eligible []Vote
	var ineligible []Vote

	for _, vote := range votes {


		ineligible = append(ineligible, vote)
	}
	return eligible, ineligible
}

func (v *Voting) DedupeByVoter(cls, eligibleVotes) {
	
}

func (v *Voting) CountVotes(cls, byVoter) {
	
}

func (v *Voting) DecideVotes(cls, nVoters, nValid, nInvalid) {
	
}

func (v *Voting) VerifyVoteSignature(cls, vote) {
	
}

func (v *Voting) VerifyVoteSchema(cls, vote) bool {
	
}
