package main

import (
	"fmt"
	bot "gopkg.in/telebot.v3"
	"log"
	"os"
	"os/signal"
	"simigo/cmd"
	"syscall"
	"time"
)

const (
	API_TOKEN string = "YOUR TOKEN BOT, GET FROM @BotFather"
)

func main() {
	pref := bot.Settings{
		Token:  API_TOKEN,
		Poller: &bot.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := bot.NewBot(pref)
	if err != nil {
		panic(err)
	}

	cmders := cmd.New(cmd.Config{
		Bot:     b,
		OwnerID: 123456,
	})

	b.Handle("/start", cmders.StartCmd)
	b.Handle("/help", cmders.HelpCmd)

	//On Text Handler
	b.Handle(bot.OnText, cmders.OnTextHandler)

	fmt.Println("Bot is running...")
	go func() {
		b.Start()
	}()

	//SIGTERM AND SIGINT HANDLER
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Println("Shutdown signal received, exiting...")
}
