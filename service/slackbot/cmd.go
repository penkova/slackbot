package slackbot

import (
	"github.com/apenkova/slackbot/service/structs"
	"github.com/nlopes/slack"
	"log"
	"strings"
)

type handler func(c *slack.Client, rtm *slack.RTM, ev *slack.MessageEvent, data []string) ////+

var handlers = map[string]handler{
	"chanID":  ChanIDHandler,
	"Hi":      GreetHandler,
	"hi":      GreetHandler,
	"Hey":     GreetHandler,
	"hey":     GreetHandler,
	"Hello":   GreetHandler,
	"hello":   GreetHandler,
	"Version": VersionHandler,
	"version": VersionHandler,
}

// HandleMessageEvent handles slack message event
func HandleMessageEvent(c *slack.Client, rtm *slack.RTM, ev *slack.MessageEvent, u string) {
	cmds := strings.Split(ev.Text, " ")
	var key string
	l := len(cmds)
	switch {
	case l == 1:
		key = cmds[0]
	case l >= 2:
		if cmds[0] == u {
			key = cmds[1]
			cmds = cmds[1:]
		} else {
			key = cmds[0]
		}
	default:
		return
	}
	if f, ok := handlers[key]; ok {
		f(c, rtm, ev, cmds)
		return
	}
}

// Worker for working with tasks
type Worker struct {
	Signal chan bool
	Queue  chan structs.MsgText
}

// Work for working with tasks
var Work = Worker{
	Signal: make(chan bool),
	Queue:  make(chan structs.MsgText, 2),
}

// ListenQueue listen queues for make, update ar invalidate distribution
func ListenQueue() {
	for {
		select {
		case msg := <-Work.Queue:
			func(item structs.MsgText) {
				err := Bot.Notification(item.Text)
				if err != nil {
					log.Println("Notification: ", err)
				}
			}(msg)
		}
	}
}
