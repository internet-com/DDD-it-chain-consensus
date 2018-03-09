package service

import (
	"github.com/junbeomlee/it-chain/consensus/domain/model/parliament"
	"github.com/junbeomlee/it-chain/consensus/domain/model/consensus"
	"errors"
)

func Elect(parliament parliament.Parliament) ([]*consensus.Representative, error){

	Representatives := make([]*consensus.Representative,0)

	if parliament.Leader == nil {
		return nil, errors.New("No Leader")
	}

	Representatives = append(Representatives,consensus.NewRepresentative(parliament.Leader.GetStringID()))

	for _, member := range parliament.Members{
		Representatives = append(Representatives, consensus.NewRepresentative(member.GetStringID()))
	}

	return Representatives, nil
}
