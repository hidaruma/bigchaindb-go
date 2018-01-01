package bigchaindb

import (
	"fmt"
	"net/http"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/config_utils"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/events"
	"log"
	"github.com/hyperledger/fabric/gossip/election"
	"github.com/cockroachdb/cockroach/pkg/server"
)

const Banner string = "" +
"****************************************************************************\n" +
"*                                                                          *\n" +
"*   Initialization complete. BigchainDB Server is ready and waiting.       *\n" +
"*   You can send HTTP requests via the HTTP API documented in the          *\n" +
"*   BigchainDB Server docs at:                                             *\n" +
"*    https://bigchaindb.com/http-api                                       *\n" +
"*                                                                          *\n" +
"*   Listening to client connections on: %s                      *\n" +
"*                                                                          *\n" +
"****************************************************************************\n"

type Plugin struct {
	EventTypes string
}

func (p *Plugin) Run() {

}

type Process struct {
	Name string
	Target func()
	Args func()
}

func (p *Process) Start() {

}
func StartEventsPlugins(exchange Exchange) {
	var plugins map[string]Plugin
	plugins = LoadEventsPlugins(Config.EventsPlugins)

	for name, plugin := range plugins {
		log.Println("Loading events plugin %s", name)
		var eventType

		var queue
		queue = exchange.GetSubscriberQueue(eventTypes)


	}
}

func Start() {
	log.Println("Initializing BigchainDB...")

	var exchange Exchange
	log.Println("Strating block")
	block.Start()

	log.Println("Starting voter")
	vote.Start()

	log.Println("Starting stale transaction monitor")
	stale.Start()

	log.Println("Starting election")
	election.Start(exchange.GetPublisherQueue())

	appServer := server.CreateServer(settings, logConfig)

	pWebAPI
	pWebAPI.Start()

	pWebsocketServer.Start()

	log.Println(Banner, )

	startEventsPlugins(exchange)

	exchange.Run()
}