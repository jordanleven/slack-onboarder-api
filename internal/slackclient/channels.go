package slackclient

import (
	"log"

	"github.com/slack-go/slack"
)

type Channel struct {
	ID          string
	Name        string
	IsMember    bool
	MemberCount int
	Description string
}

type ChannelId string

func getFormattedChannel(c slack.Channel) Channel {
	return Channel{
		Name:        c.Name,
		ID:          c.ID,
		IsMember:    c.IsMember,
		Description: c.Purpose.Value,
		MemberCount: c.NumMembers,
	}
}

func getFormattedChannels(c []slack.Channel) []Channel {
	var channels []Channel

	for _, ch := range c {
		fch := getFormattedChannel(ch)
		channels = append(channels, fch)
	}

	return channels
}

func (c *SlackClient) joinChannel(channel string) error {
	_, _, _, err := c.Client.JoinConversation(channel)
	log.Print(err)
	return err
}

func (c *SlackClient) JoinChannels(channels []string) error {
	for _, ch := range channels {
		err := c.joinChannel(ch)
		if err != nil {
			log.Println("Error joining channels")
			log.Print(err)
			return err
		}
	}

	return nil
}

// GetChannels returns a map of all channels
func (c *SlackClient) GetChannels() ([]Channel, error) {
	p := slack.GetConversationsParameters{
		ExcludeArchived: "true",
	}

	channels, _, err := c.Client.GetConversations(&p)
	if err != nil {
		log.Println("Error retrieving channels")
		log.Print(err)
		return nil, err
	}

	return getFormattedChannels(channels), nil
}
