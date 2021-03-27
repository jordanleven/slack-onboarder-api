package slackclient

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const SlackOAuthURL = "https://slack.com/api/oauth.v2.access"

type oAuthResponse struct {
	Success bool              `json:"ok"`
	Error   string            `json:"error"`
	User    oAuthResponseUser `json:"authed_user"`
}

type oAuthResponseUser struct {
	Token string `json:"access_token"`
	ID    string `json:"id"`
}

type Authorization struct {
	Token        string
	Name         string
	ProfileImage string
}

func init() {
	godotenv.Load()
}

func getRequestURL(code string) string {
	req, err := http.NewRequest("POST", SlackOAuthURL, nil)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("code", code)
	q.Add("client_id", os.Getenv("SLACK_CLIENT_ID"))
	q.Add("client_secret", os.Getenv("SLACK_CLIENT_SECRET"))
	req.URL.RawQuery = q.Encode()
	return req.URL.String()
}

func getFormattedResponse(body []byte) oAuthResponse {
	var r oAuthResponse

	err := json.Unmarshal(body, &r)
	if err != nil {
		log.Print("Error retrieving access token")
		log.Print(err)
	}
	return r
}

func getResponse(code string) oAuthResponse {
	url := getRequestURL(code)
	resp, err := http.Get(url)
	if err != nil {
		log.Print("Error retrieving access token")
		log.Print(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return getFormattedResponse(body)
}

// GetChannels returns a map of all channels
func GetAuthorization(code string) Authorization {
	response := getResponse(code)
	c := New(response.User.Token)
	u, _ := c.getUser(response.User.ID)

	authorization := Authorization{
		ProfileImage: u.ProfileImage,
		Name:         u.FullName,
		Token:        response.User.Token,
	}
	return authorization
}
