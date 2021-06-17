package hub


type Notification struct {
	Count int `json:"count"`
	Receiver string `json:"receiver"`
}

type NotifyHub struct {
	// Registered clients.
	clients map[*NotifyClient]bool
	// Register requests from the clients.
	Register chan *NotifyClient
	// Unregister requests from clients.
	Unregister chan *NotifyClient

	Notify chan *Notification

}
func NewNotifyHub() *NotifyHub {
	return &NotifyHub{
		Register:   make(chan *NotifyClient),
		Unregister: make(chan *NotifyClient),
		clients:    make(map[*NotifyClient]bool),
		Notify:     make(chan *Notification),
	}
}
func (h *NotifyHub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
		case notification := <-h.Notify:
			for client := range h.clients {
				if client.Id == notification.Receiver {
					select {
					case client.Send <- notification:
					default:
						close(client.Send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}
