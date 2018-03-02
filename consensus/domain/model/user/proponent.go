package user

import "github.com/junbeomlee/it-chain/consensus/domain/model/consensus"

type ProponentID struct{
	ID string
}

type Proponent struct{
	ProponentID ProponentID
}

func (p Proponent) StartConsensus(id consensus.ConsensusID,paliament consensus.Paliament,block consensus.Block) *consensus.Consensus {

	return &consensus.Consensus{
		ConsensusID:  id,
		Paliament:    paliament,
		Block:        block,
		CurrentState: new(consensus.IdleState)}
}
