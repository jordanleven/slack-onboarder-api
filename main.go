package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/jordanleven/slack-onboarder/internal/router"
)

func init() {
	godotenv.Load()
}

func startRoutes() {
	p := os.Getenv("PORT")
	c := router.RoutesConfiguration{
		Port: p,
	}
	router.Run(c)
}

func main() {
	startRoutes()
}
