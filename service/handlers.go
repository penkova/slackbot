package service

import (
	"fmt"
	"github.com/nlopes/slack"
)

// GreetHandler handles "hello" query
func GreetHandler(c *slack.Client, rtm *slack.RTM, ev *slack.MessageEvent, data []string) {
	rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Hello, <@%s>", ev.User), ev.Channel))
}
