package hub

import (
	"message-service/domain/model"
)

type MessageHub struct {
	// Registered clients.
	clients map[*MessageClient]bool
	// Inbound messages from the clients.
	Broadcast chan *model.MessageSendResponse
	// Register requests from the clients.
	Register chan *MessageClient
	// Unregister requests from clients.
	Unregister chan *MessageClient

}
func NewHub() *MessageHub {
	return &MessageHub{
		Broadcast:  make(chan *model.MessageSendResponse),
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
				if message.IsMessageRequest {
					if client.Id == message.ConversationRequest.LastMessage.MessageToId {
						select {
						case client.Send <- message:
						default:
							close(client.Send)
							delete(h.clients, client)
						}
					}
				} else {
					if client.Id == message.Conversation.LastMessage.MessageToId {
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
}
