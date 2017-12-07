package bot

import (
	"fmt"

	"github.com/nlopes/slack"
)

func SlackUsersList(api *slack.Client) {
	users, _ := api.GetUsers()
	for _, user := range users {
		fmt.Println(user.ID, user.Name, user.RealName)
	}
}

func SlackChannelsList(api *slack.Client) {
	channels, _ := api.GetChannels()
	for _, channel := range channels {
		fmt.Println(channel.ID, channel.Name)
	}
}
