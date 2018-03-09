package parliament

import "github.com/junbeomlee/it-chain/consensus/domain/model/consensus"

type Parliament struct {
	Leader  *Leader
	Members []*Member
}

func (p Parliament) IsNeedConsensus() bool{

	if len(p.Members) <= 1{
		return false
	}

	return true
}


func (p *Parliament) ValidateRepresentative(representatives []*consensus.Representative) bool{
	return true
}