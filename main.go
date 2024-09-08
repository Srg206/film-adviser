package main

import (
	"bufio"
	"film-adviser/reminder"
	"film-adviser/reminder/reminderbot"
	"film-adviser/reminder/reminderweb"
	"film-adviser/repository/postgres"
	"film-adviser/saver"
	"film-adviser/saver/saverbot"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	save_interface, remind_interface := loadconfigs() // read configs.txt and choose way to save and remind films (bot or web)

	// create and initialise way to story data
	var storage postgres.PostgresRepo
	storage.MustInit()

	// create saver which type we get from configs.txt
	var saver saver.Saver
	saver = senderfabric(save_interface)
	saver.MustInit(&storage)

	// create reminder which type we get from configs.txt
	var reminder reminder.Reminder
	reminder = receiverfabric(remind_interface)
	reminder.MustInit(&storage)

	// start saver and reminder in 2 goroutines

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		reminder.SendAnswer()
	}()
	go func() {
		defer wg.Done()
		saver.Handle()
	}()
	wg.Wait()

}

// func which return certain type of saver(bot or web )
func senderfabric(config string) saver.Saver {
	if config == "bot" {
		return saverbot.New()
	}
	return saverbot.New()
}

// func which return certain type of reminder(bot or web)
func receiverfabric(config string) reminder.Reminder {
	if config == "web" {
		return reminderweb.New()
	}
	return reminderbot.New()
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
