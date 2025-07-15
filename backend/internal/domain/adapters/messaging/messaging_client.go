package messaging

type Client interface {
	Connect() error
	Publish(topic string, msg []byte) error
	Subscribe(topic string, handler func(msg []byte)) error
	Close()
}
