package api

import (
	cs "github.com/junbeomlee/it-chain/consensus/domain/model/consensus"
	"github.com/junbeomlee/it-chain/consensus/domain/model/msg"
	"github.com/junbeomlee/it-chain/consensus/domain/model/user"
	"github.com/rs/xid"
	"github.com/junbeomlee/it-chain/consensus/domain/repository"
	"errors"
)

type ConsensusApi struct{
	copnsensusRepository repository.ConsensusRepository
	authenticationApi AuthenticationApi
	messageApi        MessageApi
}

func (cApi ConsensusApi) StartConsensus(id user.PeerID, block cs.Block, parliament cs.Parliament) error{

	//id의 자격 check
	leader,err := cApi.authenticationApi.IsLeader(id)

	if err != nil{
		//is not a leader
		return errors.New("Not a leader error")
	}

	//Paliament의 Validate Check
	valid := cApi.authenticationApi.IsValidParliment(parliament)

	if !valid{
		return errors.New("Not a vaild parliament")
	}

	if valid{

		if parliament.IsNeedConsensus() {
			consensus := leader.StartConsensus(cs.NewConsensusID(xid.New().String()),parliament, block)
			PreprepareMessage := consensus.CreatePreprepareMsg()

			cApi.copnsensusRepository.Save(consensus)
			cApi.messageApi.BroadCastPreprepareMsg(PreprepareMessage,consensus.Parliament.Members)

		}else{
			cApi.messageApi.RetureConfirmedBlock(block)
		}

	}
}

func (cApi ConsensusApi) ReceivePrepareMsg(id cs.ConsensusID, msg msg.PrepareMsg){

}

func (cApi ConsensusApi) ReceiveCommitMsg(id cs.ConsensusID, msg msg.CommitMsg){

}

func (cApi ConsensusApi) ReceivePreprepareMsg(id cs.ConsensusID, msg msg.PreprepareMsg){

}