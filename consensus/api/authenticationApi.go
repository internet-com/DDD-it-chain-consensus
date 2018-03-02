package api

import (
	"github.com/junbeomlee/it-chain/consensus/domain/model/user"
	"github.com/junbeomlee/it-chain/consensus/domain/model/consensus"
)

type AuthenticationApi struct{

}

func (aApi *AuthenticationApi) IsLeader(id user.PeerID) (consensus.Leader, error){

	return consensus.Leader{}, nil
}

func (aApi *AuthenticationApi) IsValidParliment(parliament consensus.Parliament) (bool){
	return true
}
