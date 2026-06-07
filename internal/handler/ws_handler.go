// Package handler
package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/maybeenang/pong-ping-v2/internal/network"
)

type WSHandler struct {
	hub *network.Hub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWSHandler(h *network.Hub) *WSHandler {
	return &WSHandler{
		hub: h,
	}
}

func (h *WSHandler) ServeWS(w http.ResponseWriter, r *http.Request) {

	roomID := r.URL.Query().Get("room")
	log.Printf("someone tried to access room : %s\n", roomID)

	room := h.hub.GetRoom(roomID)
	if room == nil {
		log.Printf("room %s notfound\n", roomID)
		http.Error(w, "room notfound", http.StatusNotFound)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("room %s error\n", roomID)
		// http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	err = network.RegisterClient(room, conn)
	if err != nil {
		log.Printf("room %s registration failed: %v", roomID, err)
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, err.Error()))
		conn.Close()
		return
	}
}
