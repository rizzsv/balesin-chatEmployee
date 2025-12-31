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