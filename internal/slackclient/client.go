package slackclient

import (
	"github.com/slack-go/slack"
)

type SlackClient struct {
	Client *slack.Client
}

// New returns our Slack client
func New(token string) SlackClient {
	c := slack.New(
		token,
		slack.OptionDebug(false),
	)
	return SlackClient{
		Client: c,
	}
}
