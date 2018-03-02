package api

import (
	cs "github.com/junbeomlee/it-chain/consensus/domain/model/consensus"
	"github.com/junbeomlee/it-chain/consensus/domain/model/msg"
)

type MessageApi struct{

}

func (mApi *MessageApi) BroadCastPreprepareMsg(PreprepareMsg msg.PreprepareMsg, member []cs.Member){

}

func (mApi *MessageApi) RetureConfirmedBlock(block cs.Block){

}