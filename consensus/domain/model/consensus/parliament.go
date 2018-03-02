package consensus

import (
	"github.com/junbeomlee/it-chain/consensus/domain/model/user"
)

type Leader struct{
	ID user.PeerID
}

func (leader Leader) StartConsensus(id ConsensusID,paliament Parliament,block Block) *Consensus {

	return &Consensus{
		ConsensusID:  id,
		Parliament:    paliament,
		Block:        block,
		CurrentState: new(IdleState)}
}

type Member struct{
	ID user.PeerID
}

type Parliament struct {
	Leader  Leader
	Members []Member
}

func (p Parliament) IsNeedConsensus() bool{

	if len(p.Members) <= 1{
		return false
	}

	return true
}