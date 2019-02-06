package main

import (
	"encoding/json"
	"github.com/apenkova/slackbot/service"
	"github.com/apenkova/slackbot/service/slackbot"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	errC   = make(chan error)
	BotKey Token
)

type Token struct {
	Token   string `json:"token"`
	WebHook string `json:"webhook"`
}

//reading token for my bot of token.json
func init() {
	file, err := ioutil.ReadFile("./token.json")
	if err != nil {
		log.Fatal("File doesn't exist")
	}
	if err := json.Unmarshal(file, &BotKey); err != nil {
		log.Fatal("Cannot parse token.json")
	}
}

func main() {
	slackbot.RunBot(BotKey.Token)

	service.RunTasks()

	// HTTP
	go run(service.ListenAndServeHTTP)
	go slackbot.ListenQueue()

	// Wait until complete
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	log.Println("WORK...")
endless:
	for {
		select {
		case err := <-errC:
			log.Fatalf("Error: %s", err.Error())
		case s := <-sig:
			log.Printf("Signal (%v) received, stoppingn\n", s)
			break endless
		}
	}
	log.Println("END.")
}

func run(f func() error) {
	if err := f(); err != nil {
		errC <- err
	}
}
