package events

const PoisonPill string = "POISON_PILL"

const EventTypes [string]int = {
	All:0,
	BlockValid:1,
	BlockInvalid:2,
}

type Event struct {
	Type int
	Data 
}

type Exchange struct {
	PublisherQueue
	Queues
}

/*

*/
func (ex *Exchange) GetPublisherQueue() {
	return ex.PublisherQueue	
}

func (ex *Exchange) GetSubscriberQueue(eventTypes [string]int) {
	if eventTypes == nil {
		eventTypes = EventTypes.All
	}
	var queue
	queue = Queue()
	ex.Queues[eventTypes] := append(ex.Queues[eventTypes], queue)

	return queue
}

func (ex *Exchange) Dispatch(event Event) {
	for eventType, queues := range ex.Queues {
		if eventType & event.Type {
			for _, queue := range queues {
				queue.Put(event)
			}
		}
	}
}

/*
Start the exchange
*/
func (ex *Exchange) Run() {
	for {
		event = ex.PublisherQueue.Get()
		if event == PoisonPill {
			return 
		} else {
			ex.Dispatch(event)
		}
	}	
}
