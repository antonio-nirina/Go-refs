package client

import (
	"github.com/antonio-nirina/go-refs/hub"
	"github.com/gorilla/websocket"
)

/*
	source: https://github.com/gorilla/websocket/tree/master/examples/chat
*/

type Client struct {
	hub *hub.Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}
