package hub

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

// Client is a middleman between the websocket connection and the hub.
type NotifyClient struct {
	hub *NotifyHub
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	Send chan *Notification

	Id string
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *NotifyClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case notification, ok := <-c.Send:
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

			notificationBytes, err := json.Marshal(notification)
			if err != nil {
				return
			}

			w.Write(notificationBytes)
			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(notificationBytes)
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
func ServeNotifyWs(hub *NotifyHub, w http.ResponseWriter, r *http.Request, userId string) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		log.Println(err)
		return
	}

	client := &NotifyClient{hub: hub, conn: conn, Send: make(chan *Notification), Id: userId}
	client.hub.Register <- client
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
}

