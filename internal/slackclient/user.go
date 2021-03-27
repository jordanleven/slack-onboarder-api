package slackclient

import (
	"log"

	"github.com/slack-go/slack"
)

type SlackUser = slack.User

type User struct {
	Username     string
	FullName     string
	DisplayName  string
	ProfileImage string
}

func getFormattedUser(user *slack.UserProfile) User {
	return User{
		Username:     user.DisplayName,
		FullName:     user.RealName,
		DisplayName:  user.DisplayNameNormalized,
		ProfileImage: user.Image192,
	}
}

// GetSlackUsers returns all current users in the Workspace
func (c *SlackClient) getUser(userID string) (User, error) {
	user, err := c.Client.GetUserProfile(userID, false)
	if err != nil {
		log.Println("Error getting user")
		log.Print(err)
		return User{}, err
	}
	return getFormattedUser(user), nil
}
