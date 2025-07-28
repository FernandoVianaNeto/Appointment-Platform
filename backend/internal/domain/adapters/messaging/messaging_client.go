package messaging

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

type Client interface {
	Connect() error
	Publish(topic string, msg []byte) error
	Subscribe(topic string, handler func(msg []byte)) error
	Close()
}
