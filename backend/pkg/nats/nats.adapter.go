package natsclient

import (
	"github.com/nats-io/nats.go"
)

type NatsClient struct {
	url string
	nc  *nats.Conn
}

func New(url string) *NatsClient {
	return &NatsClient{url: url}
}

func (c *NatsClient) Connect() error {
	nc, err := nats.Connect(c.url)
	if err != nil {
		return err
	}
	c.nc = nc
	return nil
}

func (c *NatsClient) Publish(topic string, msg []byte) error {
	return c.nc.Publish(topic, msg)
}

func (c *NatsClient) Subscribe(topic string, handler func(msg []byte)) error {
	_, err := c.nc.Subscribe(topic, func(m *nats.Msg) {
		handler(m.Data)
	})
	return err
}

func (c *NatsClient) Close() {
	if c.nc != nil {
		c.nc.Close()
	}
}
