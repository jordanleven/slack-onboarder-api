package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanleven/slack-onboarder/internal/slackclient"
)

func getAuthorizationToken(c *gin.Context) {
	code := c.Query("code")
	authorization := slackclient.GetAuthorization(code)
	c.JSON(200, gin.H{
		"authorization": authorization,
	})
}
