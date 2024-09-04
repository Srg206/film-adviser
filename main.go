package main

import (
	"film-adviser/receiver"
	"film-adviser/receiver/receiverbot"
	"film-adviser/receiver/receiverweb"
	"film-adviser/repository/postgres"
	"film-adviser/sender"
	"film-adviser/sender/senderbot"
	"sync"
)

func main() {
	var storage postgres.PostgresRepo

	storage.MustInit()

	var sender sender.Sender
	sender = senderfabric()
	sender.MustInit(&storage)

	var receiver receiver.Receiver
	receiver = receiverfabric()
	receiver.MustInit(&storage)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		receiver.SendAnswer()
	}()
	go func() {
		defer wg.Done()
		sender.Handle()
	}()
	wg.Wait()

}
func senderfabric() sender.Sender {
	return senderbot.New()
}

func receiverfabric() receiver.Receiver {
	if true {
		return receiverbot.New()
	} else {
		return receiverweb.New()
	}
}
