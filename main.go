package main

import (
	"film-adviser/receiver"
	"film-adviser/receiver/receiverbot"
	"film-adviser/sender"
	"film-adviser/sender/senderbot"
	"film-adviser/sender/senderweb"
)

func main() {
	var receiver receiver.Receiver
	receiver = receiverbot.New()
	receiver.MustInit()
	receiver.SendAnswer()

}
func senderfabric() sender.Sender {
	if true {
		return senderbot.New()
	} else {
		return senderweb.New()
	}
}
