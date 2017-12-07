package bot

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
	"github.com/tkitsunai/commango-slackbot/config"
	"github.com/tkitsunai/commango-slackbot/invoke"
)

type CommangoApi struct {
	*slack.Client
}

func NewCommangoApi(client *slack.Client) *CommangoApi {
	return &CommangoApi{
		Client: client,
	}
}

func (api *CommangoApi) SendCommandResult(commandArg string) error {
	params := slack.PostMessageParameters{
		Username:  config.Config.Slack.User,
		LinkNames: 1,
		IconEmoji: config.Config.Slack.Icon,
	}

	command := strings.Split(commandArg, " ")
	_, _, err := api.PostMessage(config.Config.Slack.Channel, invoke.CallCommand(command).String(), params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}

func (api *CommangoApi) MyEventMessage(text string) {
	isMe := strings.HasPrefix(text, config.Config.Slack.BotMention())
	if !isMe {
		return
	}
	api.SendCommandResult(strings.TrimSpace(strings.Replace(text, config.Config.Slack.BotMention(), "", -1)))
}
