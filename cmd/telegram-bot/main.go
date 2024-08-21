package main

import (
	"flag"
	"log"
	tgClient "telegram-bot/clients/telegram"
	event_consumer "telegram-bot/consumer/event-consumer"
	"telegram-bot/events/telegram"
	"telegram-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "file_storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, MustToken()),
		files.New(storagePath),
	)

	log.Print("service started")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func MustToken() string {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("token is not specified:", r)
		}
	}()

	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
