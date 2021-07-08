package tokenmanager

import (
	"fmt"
	"time"

	"github.com/jordanleven/slack-onboarder/internal/slackclient"
	uuid "github.com/nu7hatch/gouuid"
)

const TOKEN_TIME_TO_LIVE_IN_MINUTES = 30

type storedToken struct {
	SlackToken slackclient.Authorization
	Expiration time.Time
}

var storedTokenMap map[string]*storedToken

func init() {
	storedTokenMap = make(map[string]*storedToken)
}

func getTimestampCurrent() time.Time {
	return time.Now()
}

func getTokenExpiration() time.Time {
	dt := getTimestampCurrent()
	return dt.Add(time.Minute * TOKEN_TIME_TO_LIVE_IN_MINUTES)
}

func getClientToken() (*uuid.UUID, error) {
	return uuid.NewV4()
}

func SetToken(slackAuthToken slackclient.Authorization) string {
	cT, err := getClientToken()

	if err != nil {
		fmt.Print("Error generating client token")
		return ""
	}

	cTString := cT.String()
	ex := getTokenExpiration()

	storedTokenMap[cTString] = &storedToken{
		SlackToken: slackAuthToken,
		Expiration: ex,
	}

	return cTString
}
