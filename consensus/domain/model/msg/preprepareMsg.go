package msg

import (
	"github.com/junbeomlee/it-chain/consensus/domain/model/consensus"
	"github.com/junbeomlee/it-chain/common"
	"github.com/junbeomlee/it-chain/protos"
	"github.com/golang/protobuf/proto"
)

type PreprepareMsg struct {
	Consensus       consensus.Consensus
}

func (pm PreprepareMsg) ToByte() ([]byte,error){
	data, err := common.Serialize(pm)

	if err != nil{
		return nil, err
	}

	streamMsg := &protos.StreamMsg{}
	streamMsg.Content = &protos.StreamMsg_PreprepareMessage{
		PreprepareMessage:&protos.PreprepareMessage{Data:data}}

	streamData,err := proto.Marshal(streamMsg)

	if err != nil{
		return nil, err
	}

	return streamData, nil
}