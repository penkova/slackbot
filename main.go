package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/penkova/slackbot/slackbot"
)

var (
	BotKey Token
)

type Token struct {
	Token string `json:"token"`
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
	slackbot.Run(BotKey.Token)
}
