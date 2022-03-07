package hub

import (
	"github.com/antonio-nirina/go-refs/client"
)

/*
	source: https://github.com/gorilla/websocket/tree/master/examples/chat
*/
type Hub struct {
	// Registered clients.
	clients map[*client.Client]bool
	// Inbound messages from the clients.
	broadcast chan []byte
	// Register requests from the clients.
	register chan *client.Client
	// Unregister requests from clients.
	unregister chan *client.Client
}
