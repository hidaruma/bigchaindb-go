package bigchaindb

import (
	"github.com/BurntSushi/wingo/event"
	"net"
)

const POISONPILL string = "POISON_PILL"

var EventTypes = map[string]int{
	"ALL":0,
	"BLOCKVALID":1,
	"BLOCKINVALID":2,
}

type Event struct {
	Type int
	Data map[sting]interface{}
}

func (e *Event) init(eventType int, eventData map[string]string) {
	e.Type = eventType
	e.Data = eventData
}

type Queue []interface{}

type Exchange struct {
	PublisherQueue *Queue
	Queues map[int][]*Queue
}

/*

*/
func (ex *Exchange) GetPublisherQueue() *Queue {
	return ex.PublisherQueue	
}

func (ex *Exchange) GetSubscriberQueue(eventTypes int) *Queue {
	if eventTypes == nil {
		eventTypes = EventTypes["ALL"]
	}
	var queue *Queue
	ex.Queues[eventTypes] = append(ex.Queues[eventTypes], queue)

	return queue
}

func (ex *Exchange) Dispatch(event interface{}) {
	switch event.(type) {
	case *Event:
		for eventType, queues := range ex.Queues {
			if eventType == event.Type {
				for _, queue := range queues {
					*queue = append(*queue, event)
				}
			}
		}
	}
}

/*
Start the exchange
*/
func (ex *Exchange) Run() {
	for {
		event := []interface{}(*ex.PublisherQueue)
		if event[0] == POISONPILL {
			return 
		} else {
			switch event[0].(type) {
			case *Event:
				ex.Dispatch(event[0])
			}
		}
	}	
}
