package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type gitUser struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         string `json:"site_admin"`
}

func main() {
	resp, err := http.Get("https://api.github.com/users")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)

	var users []gitUser

	_ = json.Unmarshal(body, &users)

	for _, user := range users {
		fmt.Printf("User Login: %v --- User Url: %v\n", user.Login, user.Url)
	}
}
