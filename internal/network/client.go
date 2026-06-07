// Package network
package network

import (
	"errors"
	"log"

	"github.com/gorilla/websocket"
)

var errRoomFull = errors.New("room is full")

type Client struct {
	Room     *Room
	Conn     *websocket.Conn
	Send     chan []byte
	PlayerID int
}

type PlayerInput struct {
	Direction string `json:"direction"`
}

func (c *Client) readPump() {
	defer func() {
		c.Room.Unregister <- c
	}()

	for {
		var input PlayerInput
		if err := c.Conn.ReadJSON(&input); err != nil {
			log.Printf("readPump player %d room %s: %v", c.PlayerID, c.Room.ID, err)
			break
		}
		c.Room.Engine.MovePaddle(c.PlayerID, input.Direction)
	}
}

func (c *Client) writePump() {
	defer c.Conn.Close()

	for {
		message, ok := <-c.Send
		if !ok {
			// channel ditutup room — kirim close frame
			c.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			return
		}
		if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("writePump player %d room %s: %v", c.PlayerID, c.Room.ID, err)
			return
		}
	}
}

func RegisterClient(room *Room, conn *websocket.Conn) error {
	client := &Client{
		Room: room,
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	resp := make(chan registerResponse, 1)
	room.Register <- RegisterRequest{Client: client, Response: resp}

	result := <-resp
	if result.Err != nil {
		return result.Err
	}

	go client.writePump()
	go client.readPump()

	return nil
}
