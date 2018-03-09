package api

import (
	"testing"
	m "github.com/junbeomlee/it-chain/messaging"
	"github.com/junbeomlee/it-chain/consensus/domain/model/msg"
	"github.com/junbeomlee/it-chain/consensus/domain/model/consensus"
	"sync"
	"github.com/junbeomlee/it-chain/messaging/event"
	"fmt"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestMessageApi_BroadCastMsg(t *testing.T) {

	messaging := m.NewMessaging("amqp://guest:guest@localhost:5672/")
	messaging.Start()

	wg := sync.WaitGroup{}
	wg.Add(1)

	msgs, _ := messaging.Consume(event.MessageCreated.String())

	go func (){
		fmt.Println("waiting")
		for data := range msgs{
			ReceivedMsg := &m.Sendable{}
			json.Unmarshal(data.Body, ReceivedMsg)
			assert.Equal(t,[]string{"1","2"},ReceivedMsg.Ids)
			wg.Done()
		}
	}()

	mApi := NewMessageApi(messaging.Publish)

	message := msg.PreprepareMsg{}
	representatives := []*consensus.Representative{&consensus.Representative{"1"},&consensus.Representative{"2"}}

	mApi.BroadCastMsg(message,representatives)

	wg.Wait()
}
