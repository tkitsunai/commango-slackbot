package bot

import (
	"log"
	"os"

	"github.com/nlopes/slack"
	"github.com/tkitsunai/commango-slackbot/config"
)

func SlackBot() {
	token := config.Config.Slack.Token
	client := slack.New(token)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	client.SetDebug(false)

	rtm := client.NewRTM()
	go rtm.ManageConnection()

	commango := NewCommangoApi(client)
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			commango.MyEventMessage(ev.Text)
			return
		default:
		}
	}
}
