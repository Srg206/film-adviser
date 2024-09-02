package sender

// Sender is component to write saved films into repository

type Sender interface {
	MustInit()
	Handle() error
}
