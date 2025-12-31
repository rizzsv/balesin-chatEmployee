package websocket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ID string
	hub *Hub
	Conn *websocket.Conn
	Send chan []byte
}

type IncomingMessage struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

type OutgoingMessage struct {
	From    string `json:"from"`
	Message string `json:"message"`
}