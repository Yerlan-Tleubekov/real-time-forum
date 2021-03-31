package repository

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type ClientRepository struct {
	db *sql.DB
}

type Client struct {
	UserID int

	Hub *Hub

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
	Db   *sql.DB
}

type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

type ChatHubs struct {
	Hubs map[int]*Hub
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{db}
}

func NewChatHubs() *ChatHubs {
	return &ChatHubs{
		Hubs: make(map[int]*Hub),
	}
}

func (chatH *ChatHubs) Register(dialogRoomID int, hub *Hub) bool {

	_, ok := chatH.Hubs[dialogRoomID]
	if !ok {
		chatH.Hubs[dialogRoomID] = hub
		return true
	}

	return true
}

func (chatH *ChatHubs) Update(dialogRoomID int, hub *Hub) bool {
	chatH.Hubs[dialogRoomID] = hub

	return true
}

func (chatH *ChatHubs) Delete(dialogRoomID int) bool {
	delete(chatH.Hubs, dialogRoomID)

	return true
}

func (chatH *ChatHubs) GetHub(dialogRoomID int) (*Hub, bool) {
	hub, ok := chatH.Hubs[dialogRoomID]
	if !ok {
		return nil, false
	}

	return hub, true

}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (c *ClientRepository) NewClient(userID int, hub *Hub, conn *websocket.Conn, send chan []byte) *Client {
	return &Client{UserID: userID, Hub: hub, Conn: conn, Send: send, Db: c.db}
}

func (h *Hub) Run() {

	for {

		select {

		case client := <-h.Register:
			h.Clients[client] = true

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

		case message := <-h.Broadcast:

			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

func (c *Client) ReadPump(userID, dialogRoomID int) {

	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {

		_, message, err := c.Conn.ReadMessage()

		if err != nil {

			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Printf("error: %v", err)
			}

			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.SaveMessage(message, userID, dialogRoomID)

		c.Hub.Broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {

	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	// for i := 0; i <  {

	// 	break
	// }

	for {

		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			// Add queued chat messages to the current websocket message.
			n := len(c.Send)

			for i := 0; i < n; i++ {
				newMessage := <-c.Send
				w.Write(newMessage)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) SaveMessage(message []byte, userID, dialogRoomID int) {
	if _, err := c.Db.Exec(`
			INSERT INTO message_history ( dialog_room_id, message, user_id, created_date )
			VALUES (?, ?, ?, ?)
`, dialogRoomID, string(message), userID, time.Now()); err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
