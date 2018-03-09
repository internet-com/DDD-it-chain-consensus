package service

import (
	cs "github.com/junbeomlee/it-chain/consensus/domain/model/consensus"
	"github.com/junbeomlee/it-chain/consensus/domain/model/msg"
)

func CheckPreparePolicy(consensus cs.Consensus,msgPool msg.MsgPool) bool{
	//NumberOfRepresentatives := len(consensus.Representatives)
	//
	//msgPool.FindPrepareMsgsByConsensusID()
	return true
}