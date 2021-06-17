package hub

type Message struct {
	Text     string `json:"text"`
	Receiver string `json:"receiver"`
}


type MessageHub struct {
	// Registered clients.
	clients map[*MessageClient]bool
	// Inbound messages from the clients.
	Broadcast chan *Message
	// Register requests from the clients.
	Register chan *MessageClient
	// Unregister requests from clients.
	Unregister chan *MessageClient

}
func NewHub() *MessageHub {
	return &MessageHub{
		Broadcast:  make(chan *Message),
		Register:   make(chan *MessageClient),
		Unregister: make(chan *MessageClient),
		clients:    make(map[*MessageClient]bool),
	}
}
func (h *MessageHub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.clients {
				if client.Id == message.Receiver {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}
