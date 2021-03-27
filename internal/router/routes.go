package router

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type RoutesConfiguration struct {
	Port string
}

type router struct {
	Engine *gin.Engine
}

func init() {
	godotenv.Load()
}

func getAuthorization(c *gin.Context) string {
	return c.Request.Header.Get("Authorization")
}

func (r *router) setCORS() {
	appURL := os.Getenv("APP_URL")
	r.Engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{appURL},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))
}

func Run(c RoutesConfiguration) {
	r := router{
		Engine: gin.Default(),
	}

	r.setCORS()

	r.Engine.GET("/authorizationToken", getAuthorizationToken)
	r.Engine.GET("/channels", getChannels)
	r.Engine.POST("/channels", joinChannels)

	p := fmt.Sprintf(":%s", c.Port)
	r.Engine.Run(p)
}
