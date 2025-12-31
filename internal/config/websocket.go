package config

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var WSUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: Restrict in production
	},
}

func GetWSUpgrader() *websocket.Upgrader {
	return &WSUpgrader
}
