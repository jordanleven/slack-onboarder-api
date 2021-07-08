package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordanleven/slack-onboarder/internal/slackclient"
)

func getChannels(c *gin.Context) {
	b := getAuthorization(c)
	client := slackclient.New(b)
	channels, err := client.GetChannels()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error retrieving channels")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"channels": channels,
		})
	}
}
