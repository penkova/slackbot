package slackbot

import (
	"fmt"
	"github.com/nlopes/slack"
)

// ChanIDHandler handles "chanID" query
func ChanIDHandler(c *slack.Client, rtm *slack.RTM, ev *slack.MessageEvent, data []string) {
	rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Hello, <@%s>.\nChanID=%s", ev.User, ev.Channel), ev.Channel))
}

// GreetHandler handles "hello" query
func GreetHandler(c *slack.Client, rtm *slack.RTM, ev *slack.MessageEvent, data []string) {
	rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Hello, <@%s>", ev.User), ev.Channel))
}

// VersionHandler handles "version" query
func VersionHandler(c *slack.Client, rtm *slack.RTM, ev *slack.MessageEvent, data []string) {
	rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("Version: %s\nBuild: %s", "1.1", "06.2.2019"), ev.Channel))
}
