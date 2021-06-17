package hub

import (
	"fmt"
	"github.com/beevik/guid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
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
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
// Client is a middleman between the websocket connection and the hub.
type MessageClient struct {
	hub *MessageHub
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	Send chan *Message

	Id string
}
// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *MessageClient) readPump() {
	defer func() {
		c.hub.Unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		message := Message{}
		err := c.conn.ReadJSON(&message)

		fmt.Println("Message ", message.Text)
		fmt.Println("Receiver ", message.Receiver)

		//		err := conn.ReadJSON(&mm)
		//
		//		if err != nil {
		//			fmt.Println("Error reading json.", err)
		//		}
		//_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.Broadcast <- &message
	}
}
// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *MessageClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write([]byte(message.Text))
			fmt.Println("Write text ", message.Text)
			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write([]byte((<-c.Send).Text))
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
// serveWs handles websocket requests from the peer.
func ServeWs(hub *MessageHub, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		fmt.Println("odje")
		log.Println(err)
		return
	}
	id := guid.New().String()
	fmt.Println("Id ", id)
	client := &MessageClient{hub: hub, conn: conn, Send: make(chan *Message), Id: id}
	client.hub.Register <- client
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

