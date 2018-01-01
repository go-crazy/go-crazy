/**
 * 
 * File: hub.go
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 3, 2017-12-20 7:47:31 pm
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 3, 2017-12-20 8:27:09 pm
 * -----
 * Copyright 2017 - 2027 乐编程, 乐编程
 */


package WsChat

import (
	// . "github.com/go-crazy/go-crazy/Config"
)

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				// Logger.Info("Close a websocket connsction!")
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
					// Logger.Info("Close a websocket connsction!")
				}
			}
		}
	}
}
