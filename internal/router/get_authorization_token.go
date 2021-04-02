package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordanleven/slack-onboarder/internal/slackclient"
)

func getAuthorizationToken(c *gin.Context) {
	code := c.Query("code")
	authorization := slackclient.GetAuthorization(code)
	c.JSON(http.StatusOK, gin.H{
		"authorization": authorization,
	})
}
