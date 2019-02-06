package slackbot

import (
	"errors"
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"time"
)

// Slack bot errors
var (
	ErrEmptyToken  = errors.New("token was not provided")
	ErrInvalidAuth = errors.New("auth error")
	ErrConnection  = errors.New("connection error")

	Bot                 BotClient
	BotMainChan         = ""
	BotNotificationChan = ""
)

// BotClient describes Bot Client service
type BotClient interface {
	SendMessage(msg, chID string)
	Notification(msg string) (err error)
}

// RunBot runs slack bot with configure token
func RunBot(token string) {
	c := New(token)
	Bot = c
	go func() {
		if err := c.Run(); err != nil {
			log.Println("Error: run bot err:", err)
		}
	}()
	time.Sleep(time.Second * 5)
	c.SendMessage("I'm online :)", BotMainChan)
}

// SlackClient describes Slack Client Entity
type SlackClient struct {
	Token  string
	Client *slack.Client
	Rtm    *slack.RTM
}

// New returns *SlackClient
func New(token string) *SlackClient {
	return &SlackClient{Token: token}
}

// Run runs Slack Bot ------------------------------///+
func (c *SlackClient) Run() error {
	if len(c.Token) == 0 {
		return ErrEmptyToken
	}
	c.Client = slack.New(c.Token)

	c.Rtm = c.Client.NewRTM()
	err := make(chan error)
	go c.serveEvents(err)
	go c.Rtm.ManageConnection()
	return <-err
}

// SendMessage sends msg to channel
func (c *SlackClient) SendMessage(msg, chID string) {
	c.Rtm.SendMessage(c.Rtm.NewOutgoingMessage(msg, chID))
}

func (c *SlackClient) Notification(msg string) (err error) {
	//c.Rtm.SendMessage(c.Rtm.NewOutgoingMessage(msg, chID))
	c.SendMessage(msg, BotNotificationChan)
	return
}

func (c *SlackClient) serveEvents(err chan error) {
	var currentUser string
	for msg := range c.Rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			currentUser = fmt.Sprintf("<@%s>", ev.Info.User.ID)
		case *slack.HelloEvent:
		case *slack.InvalidAuthEvent:
			err <- ErrInvalidAuth
		case *slack.ConnectionErrorEvent:
			err <- ErrConnection
		case *slack.MessageEvent:
			log.Printf("MessageEvent User: %+v Text: %+v Chanel: %+v\n", ev.Msg.User, ev.Msg.Text, ev.Channel)
			HandleMessageEvent(c.Client, c.Rtm, ev, currentUser)
		}
	}
}
