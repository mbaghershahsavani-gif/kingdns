package controller

type Message struct {
	Type    string
	Payload string
}

type MessageBroker struct {
	Connected bool
}

func NewMessageBroker() *MessageBroker {
	return &MessageBroker{Connected: true}
}

func (b *MessageBroker) Publish(message Message) bool {
	return b.Connected && message.Type != ""
}
