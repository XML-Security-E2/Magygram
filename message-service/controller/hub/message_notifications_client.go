package hub

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// Client is a middleman between the websocket connection and the hub.
type MessageNotificationClient struct {
	hub *MessageNotificationsHub
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	SendNotification chan *Notification

	Id string
}

func (c *MessageNotificationClient) writeNotificationsPump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.SendNotification:
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

			messageBytes, err := json.Marshal(message)
			if err != nil {
				return
			}

			w.Write(messageBytes)
			// Add queued chat messages to the current websocket message.
			n := len(c.SendNotification)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(messageBytes)
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


func ServeMessageNotificationWs(hub *MessageNotificationsHub, w http.ResponseWriter, r *http.Request, userId string, unviewedNotifications int) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		return
	}
	client := &MessageNotificationClient{hub: hub, conn: conn, SendNotification: make(chan *Notification), Id: userId}
	client.hub.Register <- client
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writeNotificationsPump()

	client.hub.Notify <- &Notification{
		Count:    unviewedNotifications,
		Receiver: userId,
	}
}


