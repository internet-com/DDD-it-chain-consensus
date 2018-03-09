package repository

import "github.com/junbeomlee/it-chain/consensus/domain/model/consensus"

type ConsensusRepository interface{
	Save(consensus consensus.Consensus)
	FindById(consensusId consensus.ConsensusID) *consensus.Consensus
}
