package sender

type Sender interface {
	Start() error
	Handle() error
}
