package repository

import (
	"github.com/junbeomlee/it-chain/consensus/domain/model/parliament"
	"sync"
	"errors"
)

type ParlimentRepository interface {
	Get() *parliament.Parliament
	Insert(*parliament.Parliament) error
}

type ParlimentRepository_impl struct{
	lock sync.Mutex
	parliament *parliament.Parliament
}

func (pr *ParlimentRepository_impl) Get() *parliament.Parliament{
	pr.lock.Lock()
	defer pr.lock.Unlock()

	return pr.parliament
}

func (pr *ParlimentRepository_impl) Insert(parliament *parliament.Parliament) error {
	pr.lock.Lock()
	defer pr.lock.Unlock()

	if parliament == nil{
		return errors.New("nil parliament")
	}

	if parliament.Leader == nil{
		return errors.New("need leader")
	}

	pr.parliament = parliament

	return nil
}