package main

import (
	"bufio"
	"film-adviser/receiver"
	"film-adviser/receiver/receiverbot"
	"film-adviser/receiver/receiverweb"
	"film-adviser/repository/postgres"
	"film-adviser/sender"
	"film-adviser/sender/senderbot"
	"film-adviser/sender/senderweb"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	send_interface, receive_interface := loadconfigs()

	var storage postgres.PostgresRepo

	storage.MustInit()

	var sender sender.Sender
	sender = senderfabric(send_interface)
	sender.MustInit(&storage)

	var receiver receiver.Receiver
	receiver = receiverfabric(receive_interface)
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
func senderfabric(config string) sender.Sender {
	if config == "bot" {
		return senderbot.New()
	} else {
		return senderweb.New()
	}
}

func receiverfabric(config string) receiver.Receiver {
	if config == "web" {
		return receiverweb.New()
	} else {
		return receiverbot.New()
	}
}

func loadconfigs() (string, string) {

	file, err := os.Open("configs.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)

	var res []string
	for scanner.Scan() {
		res = append(res, strings.Split(scanner.Text(), " ")[1])
	}
	return res[0], res[1]
}
