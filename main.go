package main

import (
	"film-adviser/sender"
	"film-adviser/sender/telegram"
	"film-adviser/sender/web"
	"film-adviser/settings"
	"fmt"
)

func main() {

	fmt.Println("From singleton token      " + settings.GetSettings().TgSenderToken)
	var sender sender.Sender
	sender = senderfabric()
	sender.MustInit()

	sender.Handle()

}
func senderfabric() sender.Sender {
	if true {
		return telegram.New()
	} else {
		return web.New()
	}
}
