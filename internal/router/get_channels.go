package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanleven/slack-onboarder/internal/slackclient"
)

func getChannels(c *gin.Context) {
	b := getAuthorization(c)
	client := slackclient.New(b)
	channels, err := client.GetChannels()

	if err != nil {
		c.JSON(300, "Error retrieving channels")
	} else {
		c.JSON(200, gin.H{
			"channels": channels,
		})
	}
}
