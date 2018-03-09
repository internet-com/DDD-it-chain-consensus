package msg

import (
	"github.com/junbeomlee/it-chain/consensus/domain/model/consensus"
	"github.com/junbeomlee/it-chain/protos"
	"github.com/junbeomlee/it-chain/common"
	"github.com/golang/protobuf/proto"
)

type CommitMsg struct {
	ConsensusID  consensus.ConsensusID
}

func (c CommitMsg) ToByte() ([]byte,error){
	data, err := common.Serialize(c)

	if err != nil{
		return nil, err
	}

	streamMsg := &protos.StreamMsg{}
	streamMsg.Content = &protos.StreamMsg_CommitMessage{
		CommitMessage:&protos.CommitMessage{Data:data}}

	streamData,err := proto.Marshal(streamMsg)

	if err != nil{
		return nil, err
	}

	return streamData, nil
}