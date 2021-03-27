package router

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jordanleven/slack-onboarder/internal/slackclient"
)

func getChannelIDs(c *gin.Context) []string {
	var channels string

	var json struct {
		Channels string `json:"channels" binding:"required"`
	}

	if c.Bind(&json) == nil {
		channels = json.Channels
	}

	return strings.Split(channels, ",")
}

func joinChannels(c *gin.Context) {
	b := getAuthorization(c)
	chIDs := getChannelIDs(c)
	client := slackclient.New(b)
	err := client.JoinChannels(chIDs)

	if err != nil {
		c.JSON(300, "Error joining channels")
	} else {
		c.JSON(200, gin.H{
			"channels": chIDs,
		})
	}
}
