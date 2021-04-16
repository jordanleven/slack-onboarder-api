package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordanleven/slack-onboarder/internal/slackclient"
	"github.com/jordanleven/slack-onboarder/internal/tokenmanager"
)

func getAuthorizationToken(c *gin.Context) {
	code := c.Query("code")
	authorization := slackclient.GetAuthorization(code)
	token := tokenmanager.SetToken(authorization)

	c.JSON(http.StatusOK, gin.H{
		"authorization": token,
	})
}
