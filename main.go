package main

import (
	"film-adviser/receiver"
	"film-adviser/receiver/receiverbot"
	"film-adviser/receiver/receiverweb"
	"film-adviser/repository/postgres"
	"film-adviser/sender"
	"film-adviser/sender/senderbot"
	"film-adviser/sender/senderweb"
)

func main() {
	var storage postgres.PostgresRepo

	storage.MustInit()

	storagePtr := &storage // Получаем указатель на структуру

	var sender sender.Sender
	sender = senderfabric()
	sender.MustInit(storagePtr)
	sender.Handle()
}
func senderfabric() sender.Sender {
	if true {
		return senderbot.New()
	} else {
		return senderweb.New()
	}
}

func receiverfabric() receiver.Receiver {
	if true {
		return receiverbot.New()
	} else {
		return receiverweb.New()
	}
}
