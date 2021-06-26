package hub

type MessageNotificationsHub struct {
	// Registered clients.
	clients map[*MessageNotificationClient]bool
	// Register requests from the clients.
	Register chan *MessageNotificationClient
	// Unregister requests from clients.
	Unregister chan *MessageNotificationClient

	Notify chan *Notification
}
func NewMessageNotificationsHub() *MessageNotificationsHub {
	return &MessageNotificationsHub{
		Register:   make(chan *MessageNotificationClient),
		Unregister: make(chan *MessageNotificationClient),
		clients:    make(map[*MessageNotificationClient]bool),
		Notify:     make(chan *Notification),
	}
}
func (h *MessageNotificationsHub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.SendNotification)
			}

		case notification := <-h.Notify:
			for client := range h.clients {
				if client.Id == notification.Receiver {
					select {
					case client.SendNotification <- notification:
					default:
						close(client.SendNotification)
						delete(h.clients, client)
					}
				}
			}
		}

	}
}
